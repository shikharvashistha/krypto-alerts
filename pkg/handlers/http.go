package handlers

import "github.com/gin-gonic/gin"

func RegisterHTTPHandlers(e *gin.RouterGroup, app Application) {
	ep := NewEndpoint(app)
	e.POST(CreateAlert, ep.CreateAlert)   // This is the route that will be called when a user wants to create an alert
	e.GET(GetAlerts, ep.GetAlerts)        // This is the route that will be called when a user wants to get all the alerts
	e.DELETE(DeleteAlert, ep.DeleteAlert) // This is the route that will be called when a user wants to delete an alert
	e.GET(GetToken, app.GetToken)         // This is the route that will be called when a user wants to get a JWT token
}
