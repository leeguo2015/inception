package handle

import (
	"inception/api/internal/global"
	"inception/api/internal/logic/users"
	"inception/api/internal/response"
	"inception/api/internal/utils"

	"github.com/gin-gonic/gin"
)

// Login
//
//	@Description:
//	@param c
func Login(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	// todo 暂时关闭验证码
	//captcha := c.Param("captcha")
	//captchaID := c.Param("captchaID")
	//if !users.CheckCaptcha(captchaID, captcha) {
	//	global.Log.Errorf("%s:验证码错误", userName)
	//	response.HttpBadRequest("验证码错误", c)
	//}
	if userName == "" || password == "" {
		response.HttpBadRequest(utils.ErrParams, c)
		return
	}
	userInfo, token, err := users.Login(userName, password)
	if err != nil {
		global.Log.Errorf("%s 登录失败:%s,", userName, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"userInfo": userInfo,
		"token":    token,
	}, "登录成功", c)
}
