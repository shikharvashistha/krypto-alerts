package handlers

import "github.com/gin-gonic/gin"

type Application interface {
	CreateAlert(c *gin.Context) // CreateAlert implements Application to create an alert
	GetAlerts(c *gin.Context)   // GetAlerts implements Application to get all the alerts
	DeleteAlert(c *gin.Context) // DeleteAlert implements Application to delete an alert
	GetToken(c *gin.Context)    // GetToken implements Application to get a JWT token
}
