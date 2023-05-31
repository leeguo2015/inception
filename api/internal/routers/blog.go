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
	route.POST("/", handle.BlogAdd)
	route.GET("/:blogID", handle.BlogGet)
	route.DELETE("/:blogID", handle.BlogDelete)
	route.PUT("/:blogID", handle.BlogUpdate)
}
