package internal

import (
	"github.com/gin-gonic/gin"
)

func paddingRouterV1(engine *gin.Engine) {
	_ = engine.Group("/v1")
	// routers.User(V1)
}
