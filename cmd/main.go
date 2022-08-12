package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shikharvashistha/krypto-alerts/pkg/handlers"
	"github.com/shikharvashistha/krypto-alerts/pkg/store"
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"github.com/shikharvashistha/krypto-alerts/pkg/utils"
)

var port = "0.0.0.0:8080"

func main() {
	// Initialize the logger
	logger := utils.NewLogger("main")

	// Connect to Redis
	logger.Info("Attempting to connect to key value store")

	utils.RedisConnect()

	logger.Info("Attempting to connect to key value store")

	// Open connection to the database
	db := utils.GetDB()

	logger.Info("Successfully connected to the database")

	logger.Info("Attempting to register the schemas...")

	// Register the schemas
	err := models.RegisterSchema(db)
	if err != nil {
		logger.WithError(utils.ADB, err).Info("Failed to register the schemas")
	}

	logger.Info("Successfully registered the schemas")

	// Initialize the router
	r := gin.Default()
	// Group the routes
	v1 := r.Group("/v1", gin.Logger())

	// Initialize the store
	store := store.NewStore(db)

	// Initialize the App service
	appSvc := handlers.NewAppSvc(store, logger)

	// Initialize the handlers
	handlers.RegisterHTTPHandlers(v1, appSvc)

	s := &http.Server{
		Addr:         port,
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
		Handler:      r,
	}

	logger.Info("Starting the server on port " + port)

	// Start the server
	if err := s.ListenAndServe(); err != nil {
		logger.WithError(utils.Application, err).Error("Failed to start the server")
	}
}
