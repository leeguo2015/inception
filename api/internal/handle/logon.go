package handle

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/global"
	"inception/api/internal/logic"
	"inception/api/internal/model"
	"inception/api/internal/response"
	"inception/api/internal/utils"
	"net/http"
)

// Logon
//
//	@Description: 注册
//	@param c
func Logon(c *gin.Context) {
	user := new(model.User)
	if err := c.BindJSON(user); err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(utils.ErrJsonBind, c)
	}
	if err := logic.Logon(user); err != nil {
		response.FailWithMessage(utils.ErrJsonBind, c)
	}
	c.JSON(http.StatusOK, nil)
}
