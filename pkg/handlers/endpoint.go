package handlers

import "github.com/gin-gonic/gin"

func NewEndpoint(app Application) *endpoint {
	return &endpoint{
		app: app,
	}
}

type endpoint struct {
	app Application
}

func (e *endpoint) CreateAlert(c *gin.Context) {
	JWTAuthentication()(c)     // JWT Token Authentication
	JWTAuthenticationVerify(c) // JWT Token Authentication Verify
	e.app.CreateAlert(c)
}

func (e *endpoint) GetAlerts(c *gin.Context) {
	JWTAuthentication()(c)
	JWTAuthenticationVerify(c)
	e.app.GetAlerts(c)
}

func (e *endpoint) DeleteAlert(c *gin.Context) {
	JWTAuthentication()(c)
	JWTAuthenticationVerify(c)
	e.app.DeleteAlert(c)
}

func (e *endpoint) GetToken(c *gin.Context) {
	e.app.GetToken(c)
}
