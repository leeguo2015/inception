package internal

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/global"
	"inception/api/internal/middleware"
	"inception/api/internal/model"
	"log"
)

func Start() {
	model.AutoMigrate()

	global.Log.Info("Starting Inception API")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuth())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	paddingRouterV1(router)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
