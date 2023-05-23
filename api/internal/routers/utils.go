package routers

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/handle"
)

func Utils(route *gin.RouterGroup) {
	Utils := route.Group("/utils")
	Utils.GET("/captcha", handle.Captcha)
	Utils.POST("/captcha", handle.VerifyCaptcha)
}
