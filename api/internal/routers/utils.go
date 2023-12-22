package routers

import (
	"inception/api/internal/handle"

	"github.com/gin-gonic/gin"
)

func Utils(route *gin.RouterGroup) {
	Utils := route.Group("/utils")
	Utils.GET("/captcha", handle.CaptchaUser)
	Utils.POST("/captcha", handle.VerifyCaptcha)
	Utils.POST("/file", handle.Uploads)

}
