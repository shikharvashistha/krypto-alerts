package handlers

import (
	"errors"
	"os"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	JWT "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	ID      string  `json:"id"`
	Expires float64 `json:"expires"`
}

// This is the middleware function that will be used to protect routes with JWT authentication
func JWTAuthentication() func(*gin.Context) error {
	return func(c *gin.Context) error {
		// Get the JWT key from environment variable
		key := []byte(os.Getenv("JWT_KEY"))
		if key == nil {
			return errors.New("JWT_KEY not found")
		} else {
			// We will create a new JWT middleware using the key previously retrieved
			authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
				Realm:            "test zone",
				SigningAlgorithm: "HS256",
				Key:              key,
				Authorizator: func(data interface{}, c *gin.Context) bool {
					if v, ok := data.(string); ok && v == "admin" {
						return true
					}
					return false
				},
				PayloadFunc: func(data interface{}) jwt.MapClaims {
					if v, ok := data.(string); ok && v == "admin" {
						return jwt.MapClaims{
							"id": v,
						}
					}
					return jwt.MapClaims{}
				},
				Unauthorized: func(c *gin.Context, code int, message string) {
					c.JSON(code, gin.H{
						"code":    code,
						"message": message,
					})
				},
				LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
					c.JSON(code, gin.H{
						"code":    code,
						"message": message,
						"time":    time,
					})
				},
				LogoutResponse: func(c *gin.Context, code int) {
					c.JSON(code, gin.H{
						"code": code,
					})
				},
				RefreshResponse: func(c *gin.Context, code int, message string, time time.Time) {
					c.JSON(code, gin.H{
						"code":    code,
						"message": message,
						"time":    time,
					})
				},
				IdentityHandler: func(*gin.Context) interface{} {
					claims := jwt.ExtractClaims(c)
					return claims["id"]
				},
				IdentityKey:   "id",
				TokenLookup:   "header: Authorization, query: token, cookie: jwt",
				TokenHeadName: "Bearer",
				TimeFunc: func() time.Time {
					return time.Now()
				},
			})
			if err != nil {
				return err
			}
			errInit := authMiddleware.MiddlewareInit()
			if errInit != nil {
				return errInit
			}
		}
		return nil
	}
}

// This is the middleware function that will be used to verify if the current JWT token is valid or nott
func JWTAuthenticationVerify(c *gin.Context) (*Token, error) {
	// Get the Bearer token from the header
	bearToken := c.Request.Header.Get("Authorization")

	token := strings.Split(bearToken, " ")
	// Retrieve the token from the request header
	if len(token) == 2 {
		parsedToken, err := JWT.Parse(token[1], func(token *JWT.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			return &Token{}, nil
		}
		// Check if the token is valid
		claims, ok := parsedToken.Claims.(JWT.MapClaims)
		if ok && parsedToken.Valid {
			expiresAt := claims["exp"].(float64)
			userID := claims["id"].(string)
			if time.Now().Unix() > int64(expiresAt) {
				return &Token{}, errors.New("Token expired")
			}
			return &Token{
				Expires: expiresAt,
				ID:      userID,
			}, nil
		}
	} else {
		return &Token{}, errors.New("invalid token")
	}
	return &Token{}, errors.New("invalid token")

}
