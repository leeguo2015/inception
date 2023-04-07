package routers

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/handle"
)

func Blog(route *gin.RouterGroup) {
	user := route.Group("/blog")
	Info(user)
}

func Info(route *gin.RouterGroup) {
	route.POST("/", handle.Login)
	route.GET("/:blogID", handle.Login)
	route.DELETE("/:blogID", handle.Login)
	route.PUT("/:blogID", handle.Login)
}
