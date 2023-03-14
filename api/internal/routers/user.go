package routers

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/handle"
)

func User(route *gin.RouterGroup) {
	user := route.Group("/user")
	Logon(user)
}

func Logon(route *gin.RouterGroup) {
	route.POST("/logon", handle.Logon)
}
