package server

import (
	"fmt"
	"interviewDemo/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "It's works!",
	})
}

func coinInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func coinAdd(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func coinUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func coinDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
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

	v2 := router.Group("/api/v2/stock")
	v2.POST("/add", coinAdd)
	v2.GET("/info", coinInfo)
	v2.PUT("/change", coinUpdate)
	v2.DELETE("/delete", coinDelete)

	if err := router.Run(httpListen); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
