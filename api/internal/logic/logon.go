package logic

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"inception/api/internal/global"
	"inception/api/internal/model"
	"time"
)

func Logon(user *model.User) error {
	if !errors.Is(global.DB.Where("username = ?", user.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	user.Uuid = uuid.New().String()
	user.Portrait = model.DefaultPortrait
	if user.Birth == "" {
		time.Now().Format(time.DateOnly)
	}
	global.DB.Create(user)
	return nil
}
