package internal

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// Simple group: v1

	paddingRouterV1(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
