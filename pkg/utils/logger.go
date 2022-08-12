package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(service string) *Logger { // NewLogger for creating a new logger

	var logger = &Logger{logrus.WithFields(logrus.Fields{"service": service})}

	return logger
}

type Type string

const (
	Application Type = "Aplication"
	ADB         Type = "DB"
	WebSocket   Type = "Websocket"
	Mail        Type = "Mail"
)

func (logger Logger) WithError(errType Type, err error) Logger { // Logger for logging errors
	return Logger{logger.WithFields(logrus.Fields{"error": err, "Type": errType})}
}

func (logger Logger) WithToken(requestToken string) Logger { // Logger for errors with token
	return Logger{logger.WithField("requestToken", requestToken)}
}

func (logger Logger) WithDBError(errType Type, err *gorm.DB) Logger { // Logger for logging database errors
	return Logger{logger.WithFields(logrus.Fields{"error": err, "Type": errType})}
}
