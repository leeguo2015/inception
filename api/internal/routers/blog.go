package routers

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/handle"
)

func Blog(route *gin.RouterGroup) {
	blogGroup := route.Group("/blog")
	Info(blogGroup)
}

// 目前权限全放开， 不考虑权限问题
func Info(route *gin.RouterGroup) {
	route.POST("/", handle.BlogAdd)
	route.GET("/:blogID", handle.BlogGet)
	route.DELETE("/:blogID", handle.BlogDelete)
	route.PUT("/:blogID", handle.BlogUpdate)
}
