package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shikharvashistha/krypto-alerts/pkg/store"
	"github.com/shikharvashistha/krypto-alerts/pkg/store/relational/models.go"
	"github.com/shikharvashistha/krypto-alerts/pkg/utils"
)

func NewAppSvc(store store.Store, logger *utils.Logger) Application {
	return &appSvc{
		Store:  store,
		Logger: logger,
	}
}

type appSvc struct {
	Store  store.Store
	Logger *utils.Logger
}

// CreateAlert implements Application
func (svc *appSvc) CreateAlert(c *gin.Context) {

	// Verify JWT token
	token, err := JWTAuthenticationVerify(c)
	if err != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}
	// Parse request body
	body := new(models.Alert)
	var price float64
	var coins []models.Coins

	if err := c.BindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// Get price of coin
	if response, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=USD&order=market_cap_desc&per_page=1&page=1&sparkline=false"); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		defer response.Body.Close()
		if err := json.NewDecoder(response.Body).Decode(&coins); err != nil {
			c.JSON(500, gin.H{
				"code":    500,
				"message": err.Error(),
			})
			return
		}
		price = coins[0].Price
	}
	// Check an entry in the database for the coin
	if err := svc.Store.RL().Alert().Create(&models.Alert{
		Model:       models.Model{UserID: token.ID},
		AlertID:     uuid.New().String(),
		Email:       body.Email,
		TriggerMail: false,
		AlertValue:  price,
		Status:      "active",
	}); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Alert created successfully",
	})

}

// DeleteAlert implements Application
func (svc *appSvc) DeleteAlert(c *gin.Context) {

	// Parse request body
	id := c.Param("id")
	alert := models.Alert{
		AlertID: id,
	}

	// Delete alert from database
	err := svc.Store.RL().Alert().Delete(&alert)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error whsile deleting alert",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Alert deleted successfully",
	})
}

// GetAlerts implements Application
func (svc *appSvc) GetAlerts(c *gin.Context) {
	var Alerts []string

	// Get the page number & page size from the query params
	pageSizeString := c.DefaultQuery("pageSize", "10")
	pageNumberString := c.DefaultQuery("pageNumber", "1")
	status := c.DefaultQuery("status", "active")

	// Convert string to int
	pageSize, err := strconv.Atoi(pageSizeString)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while converting page size to int")

	}

	// Convert string to int
	pageNumber, err := strconv.Atoi(pageNumberString)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while converting page number to int")

	}

	// Get the token data & verify the token
	token, err := JWTAuthenticationVerify(c)
	if err != nil {
		c.JSON(401, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}
	// Get the alerts from the database
	alert := models.Alert{
		Model: models.Model{UserID: token.ID},
	}
	err = svc.Store.RL().Alert().Get(&alert)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while getting alerts")
	}
	// Append the alerts to the Alerts array
	Alerts = append(Alerts, alert.AlertID)
	key := models.Cache{
		Model:    models.Model{UserID: token.ID},
		PageNo:   pageNumber,
		PageSize: pageSize,
		Status:   status,
	}
	// Get the alerts from the cache
	key_json, err := json.Marshal(key)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while marshalling key")
	}

	keys, err := svc.Store.KV().Get(string(key_json))
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while getting keys")
	}
	Alerts = append(Alerts, keys)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Alerts retrieved successfully",
		"Alerts":  Alerts,
	})

	// Paginate the alerts
	offset := pageSize * (pageNumber - 1)
	var alerts []models.Alert

	if alerts, err = svc.Store.RL().Alert().GetByOffset(&models.Alert{Model: models.Model{UserID: token.ID}}, offset, pageSize); err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while getting alerts")
	}
	var alertString string
	for _, alert := range alerts {
		alertString += alert.AlertID + ","
	}
	// Set the alerts in the cache
	timeDurationString := os.Getenv("CACHE_TIMEOUT")
	timeDuration, err := time.ParseDuration(timeDurationString)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while parsing time duration")
	}
	err = svc.Store.KV().Set(string(key_json), alertString, timeDuration)
	if err != nil {
		svc.Logger.WithError(utils.Application, err).Error("Error while setting keys")
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Alerts retrieved successfully",
		"Alerts":  alerts,
	})

}

// GetToken implements Application
func (svc *appSvc) GetToken(c *gin.Context) {

	// Get the required data from environment variables
	secret := os.Getenv("JWT_SECRET")
	timeValidity := os.Getenv("JWT_TIMEOUT")

	claims := jwt.MapClaims{
		"user_id": uuid.New().String(),
	}
	// Create a new token
	stringToTime, _ := time.ParseDuration(timeValidity)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(stringToTime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error while generating token",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": t,
	})
}
