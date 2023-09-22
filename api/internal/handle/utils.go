package handle

import (
	"inception/api/internal/global"
	"inception/api/internal/response"
	"inception/api/internal/utils/captcha"

	"github.com/gin-gonic/gin"
)

func CaptchaUser(c *gin.Context) {
	b := new(captcha.BaseApi)
	id, b64, err := b.Captcha()
	if err != nil {
		global.Log.Errorf("获取验证码失败:%s,", err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"captcha_id": id,
		"b64":        b64,
	}, "请求成功", c)
}

func VerifyCaptcha(c *gin.Context) {
	CaptchaID := c.Param("captcha_id")
	answer := c.Param("answer")
	r := captcha.NewDefaultRedisStore()
	if r.Verify(CaptchaID, answer, true) {
		response.Ok(c)
	}
}
