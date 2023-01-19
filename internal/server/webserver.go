package server

import (
	"fmt"
	"interviewDemo/internal/logger"
	"interviewDemo/internal/model"
	"interviewDemo/internal/repo"
	"interviewDemo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CoinItemRequestBody is type CoinItem for receiving through REST
type CoinItemRequestBody struct {
	SymbolID    string `json:"symbol"`
	Name        string `json:"name"`
	TypeID      uint8  `json:"type"`
	ExchangesID []uint `json:"exchanges"`
}

var coinRepo = repo.NewCoinList()
var coinService = service.NewCoinListServise(coinRepo)

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "It's works!",
	})
}

func coinInfo(c *gin.Context) {
	symbol := c.Param("symbol")
	coin, err := coinService.Read(symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "find successfully", "data": &coin})
}

func coinAdd(c *gin.Context) {
	coinBody := CoinItemRequestBody{}

	if err := c.BindJSON(&coinBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	coinItem := model.CoinItem{
		Name:        coinBody.Name,
		TypeID:      coinBody.TypeID,
		ExchangesID: coinBody.ExchangesID,
	}

	if err := coinService.Create(coinBody.SymbolID, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": true, "message": "created successfully", "data": &coinItem})
}

func coinUpdate(c *gin.Context) {
	symbol := c.Param("symbol")
	coinItem := model.CoinItem{}
	if err := c.BindJSON(&coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}

	if err := coinService.Update(symbol, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "updated successfully", "data": &coinItem})
}

func coinDelete(c *gin.Context) {
	symbol := c.Param("symbol")
	if err := coinService.Delete(symbol); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "deleted successfully", "data": &symbol})
}

func logMiddleware(c *gin.Context) {
	logger.Log.Sugar().Debugf("New visit: %s", c.Request.URL)
	c.Next()
}

// StartSRV starts web server
func StartSRV(httpListen string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Use(logMiddleware)

	router.GET("/", sayHello)

	v2 := router.Group("/api/v2/coin")
	v2.POST("/add", coinAdd)
	v2.GET("/info/:symbol", coinInfo)
	v2.PUT("/change/:symbol", coinUpdate)
	v2.DELETE("/delete/:symbol", coinDelete)

	if err := router.Run(httpListen); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
