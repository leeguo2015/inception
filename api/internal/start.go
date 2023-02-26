package internal

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	// Simple group: v1

	paddingRouterV1(router)

	router.Run(":8080")
}
