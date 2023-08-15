package routers

import (
	"inception/api/internal/handle"

	"github.com/gin-gonic/gin"
)

func Blog(route *gin.RouterGroup) {
	blogGroup := route.Group("/blog")
	CommentGroup := route.Group("/Comment")
	Info(blogGroup)
	Comment(CommentGroup)
}

// 目前权限全放开， 不考虑权限问题

func Info(route *gin.RouterGroup) {
	route.POST("/", handle.BlogAdd)
	route.GET("/:blogID", handle.BlogGet)
	route.DELETE("/:blogID", handle.BlogDelete)
	route.PUT("/:blogID", handle.BlogUpdate)
}

func Comment(route *gin.RouterGroup) {
	route.POST("/:blogID", handle.CommentAdd)
	route.GET("/:blogID", handle.CommentGet)
	route.DELETE("/:blogID", handle.CommentDelete)
	// route.PUT("/:blogID", handle.CommentUpdate)
}
