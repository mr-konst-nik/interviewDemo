package server

import (
	"fmt"
	"interviewDemo/internal/logger"
	"interviewDemo/internal/repo"
	"interviewDemo/internal/service"

	"github.com/gin-gonic/gin"

	_ "interviewDemo/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example for job interviews
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v2

var coinRepo = repo.NewCoinList()
var coinService = service.NewCoinListServise(coinRepo)

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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(httpListen); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
