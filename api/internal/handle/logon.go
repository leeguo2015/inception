package handle

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/global"
	"inception/api/internal/logic/users"
	"inception/api/internal/response"
	"inception/api/internal/utils"
)

// Logon
//
//	@Description: 注册
//	@param c
func Logon(c *gin.Context) {
	user := new(users.Register)
	if err := c.BindJSON(user); err != nil {
		global.Log.Error(err.Error())
		response.HttpUnprocessableEntity(utils.ErrJsonBind, c)
		return
	}
	if err := users.Logon(user); err != nil {
		response.HttpBadRequest(err.Error(), c)
		return
	}
	response.HttpOk("注册成功", c)
}
