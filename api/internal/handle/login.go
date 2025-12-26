package handle

import (
	"inception/api/internal/global"
	"inception/api/internal/logic/users"
	"inception/api/internal/response"
	"inception/api/internal/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login
//
//	@Description:
//	@param c
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("登录参数绑定错误：", err.Error())
		response.HttpBadRequest("参数格式错误", c)
		return
	}
	
	userName := req.Username
	password := req.Password
	
	// todo 暂时关闭验证码
	//captcha := c.Param("captcha")
	//captchaID := c.Param("captchaID")
	//if !users.CheckCaptcha(captchaID, captcha) {
	//	global.Log.Errorf("%s:验证码错误", userName)
	//	response.HttpBadRequest("验证码错误", c)
	//}
	global.Log.Info("登录信息", userName, password)
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
