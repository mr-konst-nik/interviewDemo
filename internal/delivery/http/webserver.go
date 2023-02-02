package http

import (
	"fmt"
	"interviewDemo/internal/repo"
	"interviewDemo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

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

func logMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() != http.StatusOK {
			logger.Errorf("Error status code: %v on endpoint: %s", c.Writer.Status(), c.Request.URL)
			return
		}
		logger.Debugf("New visit: %s", c.Request.URL)
	}
}

// StartSRV starts web server
func StartSRV(httpListen string, logger *zap.SugaredLogger) error {

	coinRepo := repo.NewCoinList()
	coinService := service.NewCoinListServise(coinRepo)
	h := NewHandlers(coinService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLGlob("web/templates/*")
	router.Use(logMiddleware(logger))

	router.GET("/", h.SayHello)

	v2 := router.Group("/api/v2/coin")
	v2.POST("/add", h.CoinAdd)
	v2.GET("/info/:symbol", h.CoinInfo)
	v2.PUT("/change/:symbol", h.CoinUpdate)
	v2.DELETE("/delete/:symbol", h.CoinDelete)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(httpListen); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}
	return nil
}
