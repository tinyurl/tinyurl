package server

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tinyurl/tinyurl/entity"
)

// BuildEngine return gin.Engine with route
func BuildEngine(appService *entity.ServiceProvider) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		//AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			logrus.Info("origin is ", origin)
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/health", HealthCheck)
	router.GET("/n/:shortpath", WrapeService(appService, ParseURL))
	router.POST("/api/v1/shorten", WrapeService(appService, ShortenURL))

	swaggerURL := ginSwagger.URL(appService.GlobalConfig.SwaggerURL)
	fmt.Printf("lwq %s \n", appService.GlobalConfig.SwaggerURL)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	return router
}

// Start start server
func Start(addr string, appService *entity.ServiceProvider) {
	router := BuildEngine(appService)
	router.Run(addr)
}

type RequestHandler func(*gin.Context, *entity.ServiceProvider)

func WrapeService(appService *entity.ServiceProvider, handler RequestHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, appService)
	}
}
