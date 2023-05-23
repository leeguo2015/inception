package internal

import (
	"inception/api/internal/routers"

	"github.com/gin-gonic/gin"
)

func paddingRouterV1(engine *gin.Engine) {
	V1 := engine.Group("v1")
	api := V1.Group("api")
	routers.User(api)
	routers.Blog(api)
	routers.Utils(api)
}
