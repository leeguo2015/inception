package handle

import (
	"github.com/gin-gonic/gin"
	"inception/api/internal/global"
	"inception/api/internal/logic/users"
	"inception/api/internal/response"
	"inception/api/internal/utils"
)

// Login
//
//	@Description:
//	@param c
func Login(c *gin.Context) {
	userName := c.Param("user_name")
	password := c.Param("password")
	captcha := c.Param("captcha")
	captchaID := c.Param("captchaID")
	if !users.CheckCaptcha(captchaID, captcha) {
		global.Log.Errorf("%s:验证码错误", userName)
		response.HttpBadRequest("验证码错误", c)
	}
	if userName == "" || password == "" {
		response.HttpBadRequest(utils.ErrParams, c)
	}
	userInfo, token, err := users.Login(userName, password)
	if err != nil {
		global.Log.Errorf("%s 登录失败:%s,", userName, err.Error())
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithDetailed(gin.H{
		"userInfo": userInfo,
		"token":    token,
	}, "登录成功", c)
}
