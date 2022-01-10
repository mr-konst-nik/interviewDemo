package server

import (
	"interviewDemo/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func StartSRV(httpListen string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(logMiddleware)

	router.GET("/", sayHello)
	router.POST("/stock", stockInfo)
	router.Run(httpListen)
}
