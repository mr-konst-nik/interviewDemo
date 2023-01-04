package server

import (
	"fmt"
	"interviewDemo/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Stock is type describe stocks
type Stock struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name,omitempty"`
	Exchange string `json:"exchange"`
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func stockInfo(c *gin.Context) {
	var json Stock
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol":   json.Symbol,
		"name":     json.Name,
		"exchange": json.Exchange,
	})
}

func logMiddleware(c *gin.Context) {
	logger.Log.Sugar().Debugf("New visit: %s", c.Request.URL)
	c.Next()
}

// StartSRV starts web server
func StartSRV(httpListen string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logMiddleware)

	router.GET("/", sayHello)
	router.POST("/stock", stockInfo)
	if err := router.Run(httpListen); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
