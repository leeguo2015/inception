package internal

import (
	"inception/api/internal/global"
	"inception/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	// autoMigrate()

	global.Log.Info("Starting Inception API")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.JWTAuth())
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	paddingRouterV1(router)
	if err := router.Run(":8080"); err != nil {
		global.Log.Error(err)
	}
}
