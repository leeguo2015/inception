package captcha

import (
	"github.com/mojocn/base64Captcha"
	"inception/api/internal/global"
)

// Store 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
var Store = NewDefaultRedisStore()

//var Store = base64Captcha.DefaultMemStore

type BaseApi struct{}

func (b *BaseApi) Captcha() (id, b64s string, err error) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.Captcha.ImgHeight, global.Captcha.ImgWidth, global.Captcha.KeyLong, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, Store.UseWithCtx(c)) // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, Store)
	//cp := base64Captcha.NewCaptcha(driver, Store)
	return cp.Generate()
}
