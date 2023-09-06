package users

import (
	"errors"
	"inception/api/internal/global"
	"inception/api/internal/model"
	"inception/api/internal/utils"
	"inception/api/internal/utils/captcha"
	"inception/api/internal/utils/jwt"

	"gorm.io/gorm"
)

func Login(userName, password string) (*model.User, string, error) {
	user := new(model.User)
	if global.DB.Where("user_name = ?", userName).First(user).Error == gorm.ErrRecordNotFound {
		return nil, "", errors.New("用户不存在")
	}
	// 后续添加输出错误次数
	if !utils.BcryptCheck(password, user.Password) {
		return nil, "", errors.New("密码错误")
	}
	token, err := GenerateToken(user)
	if err != nil {
		return nil, "", errors.New("生成token错误")
	}
	return user, token, nil
}

func CheckCaptcha(captchaID, captchaStr string) bool {
	return captcha.Store.Verify(captchaID, captchaStr, true)
}

func GenerateToken(user *model.User) (string, error) {
	c := jwt.BaseClaims{
		UUID:        user.Uuid,
		UserID:      user.ID,
		Username:    user.UserName,
		AuthorityId: user.AuthorityId,
	}
	return jwt.GenerateToken(c, global.JwtTimout)
}
