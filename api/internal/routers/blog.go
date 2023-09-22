package routers

import (
	"inception/api/internal/handle"

	"github.com/gin-gonic/gin"
)

func Blog(route *gin.RouterGroup) {
	blogGroup := route.Group("/blog")
	CommentGroup := route.Group("/comment")
	Info(blogGroup)
	Comment(CommentGroup)
}

// 目前权限全放开， 不考虑权限问题

func Info(route *gin.RouterGroup) {
	route.POST("/", handle.BlogAdd)
	route.GET("/:blog_id", handle.BlogGet)
	route.DELETE("/:blog_id", handle.BlogDelete)
	route.PUT("/:blog_id", handle.BlogUpdate)
}

func Comment(route *gin.RouterGroup) {
	route.POST("/:blog_id", handle.CommentAdd)
	route.GET("/:blog_id", handle.CommentGet)
	route.DELETE("/:blog_id", handle.CommentDelete)
	// route.PUT("/:blogID", handle.CommentUpdate)
}
