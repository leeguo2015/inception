package internal

import (
	"inception/api/internal/routers"

	"github.com/gin-gonic/gin"
)

func paddingRouterV1(engine *gin.Engine) {
	V1 := engine.Group("/v1")
	routers.User(V1)
	routers.Blog(V1)
}
