package internal

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Starting Inception API")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
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
