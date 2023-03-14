package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	gorm.Model
	Uuid     string    `json:"uuid" gorm:"column:uuid;type:varchar(100);"`
	UserName string    `json:"user_name" gorm:"type:varchar(32);"`
	RealName string    `json:"real_name" gorm:"type:varchar(32);"`
	Password string    `json:"password" gorm:"type:varchar(128);"`
	Portrait string    `json:"portrait" gorm:"type:varchar(128);"`
	Gender   int       `json:"gender"`
	IDCard   string    `json:"id_card" gorm:"type:varchar(128);"`
	Addr     string    `json:"addr" gorm:"type:varchar(256);"`
	Phone    string    `json:"Phone" gorm:"type:varchar(128);"`
	Email    string    `json:"email" gorm:"type:varchar(128);"`
	IPAddr   string    `json:"ip_addr" gorm:"type:varchar(128);"`
	Status   int       `json:"status"` // 0 正常，1。冻结，2.注销
	Birth    string    `json:"birth"`
	LastTime time.Time `json:"last_time"`
	//Roles    []AuthRole `gorm:"many2many:user_roles;"`
}

const (
	DefaultPortrait = "default/portrait.png"
)

func NewUser() *User {
	return &User{
		Uuid:     uuid.New().String(),
		Portrait: DefaultPortrait,
		Gender:   0,
		Birth:    time.Now().Format(time.DateOnly),
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
