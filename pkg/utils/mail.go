package utils

import (
	"net/smtp"
	"os"
)

func InitializeEmailService() {
	// Initialize the logger
	logger := NewLogger("mail")

	// Get the existing redis client or initialize a new one
	rdb, ctx := RedisConnect()

	// PubSub channel to listen for new emails
	subscription := rdb.Subscribe(ctx, "mail")

	go func() {
		for {
			// Wait for a new message
			message, err := subscription.ReceiveMessage(ctx)
			if err != nil {
				logger.WithError(Mail, err).Info("Failed to receive the message from the redis")
			}
			// Log the message
			logger.WithField("message", message).Info("Message received")
			err = sendMail([]string{message.Payload})
			if err != nil {
				logger.WithError(Mail, err).Info("Failed to send the email")
			}

			// Publish the message to the channel
			rdb.Publish(ctx, "mail", "Email sent")
			rdb.Ping(ctx)
		}
	}()
}

func sendMail(body []string) error {

	//Get the necessary environment variables for shooting the email
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("FROM_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Create the message
	message := []byte(body[0])
	// Connect to the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send the email
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{from}, message); err != nil {
		return err
	}
	return nil
}
