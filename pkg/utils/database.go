package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"
)

var (
	db          *gorm.DB      = nil
	rdb         *redis.Client = nil
	ctx                       = context.Background()
	usernameADB               = os.Getenv("USERNAME")
	passwordADB               = os.Getenv("PASSWORD")
	dbNameADB                 = os.Getenv("DBNAME")
	hostADB                   = os.Getenv("HOST")
	portADB                   = 5432

	dsnADB = fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s sslmode=disable", hostADB, usernameADB, dbNameADB, portADB, passwordADB)
)

func GetDB() *gorm.DB {
	if db == nil { // Initialize a new database connection
		var err error
		db, err = connect(dsnADB, NewLogger("main"))
		if err != nil {
			NewLogger("main").WithError(ADB, err).Info("Failed to connect to the database")
			os.Exit(1)
		}
	} else {
		return db // Return the existing database connection if it exists
	}
	return db
}
func connect(dsnURI string, logger *Logger) (*gorm.DB, error) {
	// Open connection to the database
	var (
		db      *gorm.DB
		err     error
		retries int = 10 // Number of retries to connect to the database
	)
	for retries > 0 {
		// Attempt to connect to the database
		db, err = gorm.Open(postgres.Open(dsnURI), &gorm.Config{
			SkipDefaultTransaction: false,
			Logger:                 dbLogger.Default.LogMode(dbLogger.Info),
			PrepareStmt:            true,
			FullSaveAssociations:   true,
		})
		if err != nil {
			logger.WithError(ADB, err).Error("Error connecting to the database")
			retries--
			time.Sleep(time.Second * 10) // Wait 10 seconds before retrying
		} else {
			break
		}
	}
	if err != nil {
		logger.WithError(ADB, err).Error("Error connecting to the database")
		return nil, err
	}

	postgresDB, err := db.DB() // Get the underlying database connection
	if err != nil {
		logger.WithError(ADB, err).Error("Error getting postgres DB")
		return nil, err
	}

	postgresDB.SetMaxIdleConns(10)           // Set the maximum number of connections in the idle connection pool
	postgresDB.SetMaxOpenConns(10)           // Set the maximum number of open connections to the database
	postgresDB.SetConnMaxLifetime(time.Hour) // Set the maximum lifetime of connections to the database

	return db, nil
}

func RedisConnect() (*redis.Client, context.Context) {
	// Check if the redis client is already initialized
	if rdb != nil {
		return rdb, ctx
	}
	// Initialize logger
	logger := NewLogger("cache")

	// Get the redis address and password from the environment variables
	address := os.Getenv("REDIS_ADDRESS")

	password := os.Getenv("REDIS_PASSWORD")

	// Create a new redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	// Ping the redis server for connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.WithError(ADB, err).Info("Failed to initialize the redis client")
		return nil, nil
	}
	return rdb, ctx
}
