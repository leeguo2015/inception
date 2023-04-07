package model

import (
	"errors"
	"github.com/google/uuid"
	"inception/api/internal/global"
	"time"
)

func UserMigrate() error {
	if err := global.DB.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

type User struct {
	//gorm.Model
	Base
	Uuid     uuid.UUID `json:"uuid" gorm:"index;column:uuid"`
	UserName string    `json:"userName" gorm:"index;type:varchar(32);unique_index"`
	Password string    `json:"-" gorm:"type:varchar(128);"`
	RealName string    `json:"realName" gorm:"type:varchar(32);"`
	//Portrait string    `json:"portrait" gorm:"type:varchar(128);"`
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Gender    int    `json:"gender" gorm:"type:TINYINT;default:0"`
	IDCard    string `json:"IDCard" gorm:"type:varchar(128);"`
	Addr      string `json:"addr" gorm:"type:varchar(256);"`
	Phone     string `json:"phone" gorm:"type:varchar(128);unique_index;"`
	Email     string `json:"email" gorm:"type:varchar(128);"`
	IPAddr    string `json:"IPAddr" gorm:"type:varchar(128);"`

	AuthorityId uint `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
	//Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Status   int       `json:"status" gorm:"type:int;"` // 0 正常，1。冻结，2.注销
	Birth    string    `json:"birth"`
	LastTime time.Time `json:"lastTime"`
	//Roles    []AuthRole `gorm:"many2many:user_roles;"`
	BaseTime
}

const (
	DefaultPortrait = "default/portrait.png"
)

func NewUser() *User {
	return &User{
		//Uuid:     uuid.New().String(),
		Uuid:      uuid.New(),
		HeaderImg: DefaultPortrait,
		Gender:    0,
		Birth:     time.Now().Format(time.DateOnly),
	}
}

func (u *User) Verify() error {
	if len(u.UserName) > 32 {
		return errors.New("user name too long")
	}
	if len(u.RealName) > 32 {
		return errors.New("real name too long")
	}
	if len(u.IDCard) > 128 {
		return errors.New("id card too long")
	}
	return nil
}
