package routers

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/handle"
)

func User(route *gin.RouterGroup) {
	user := route.Group("/user")
	Logon(user)
	Login(user)
}

func Logon(route *gin.RouterGroup) {
	route.POST("/logon", handle.Logon)
}

func Login(route *gin.RouterGroup) {
	route.POST("/login", handle.Login)
}
