package users

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"inception/api/internal/global"
	"inception/api/internal/model"
	"inception/api/internal/utils"
	"time"
)

// Register
// @Description:
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	HeaderImg   string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId uint   `json:"authorityId" gorm:"default:888"`
}

func Logon(register *Register) error {
	user := model.NewUser()
	if err := register.User(user); err != nil {
		return err //
	}
	if !errors.Is(global.DB.Where("user_name = ?", user.UserName).Or("email = ?", user.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名或邮箱已注册")
	}
	user.Uuid = uuid.New()
	user.HeaderImg = model.DefaultPortrait
	user.LastTime = time.Now()
	if user.Birth == "" {
		time.Now().Format(time.DateTime)
	}
	user.Password = utils.BcryptHash(user.Password)
	return global.DB.Create(&user).Error
}

func (r *Register) User(user *model.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	user.UserName = r.Username
	user.Password = r.Password
	user.HeaderImg = r.HeaderImg
	user.AuthorityId = r.AuthorityId
	return nil
}
