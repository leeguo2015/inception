package model

import (
	"inception/api/internal/global"
)

type (
	Tag struct {
		Base
		Content string `gorm:"not null;type:varchar(48)"`
		BaseTime
		Blogs []Blog `gorm:"many2many:blog_tags;"`
	}

	Blog struct {
		Base
		UUID   string
		UserID uint `gorm:"not null"`
		User   User `gorm:"foreignKey:UserID"`

		Title   string `gorm:"not null;type:varchar(48)"`
		Content string ` gorm:"not null;type:text"`
		BaseTime

		Remakes []Remake
		Likes   []Like
		Tags    []Tag `gorm:"many2many:blog_tags;"`
	}

	Remake struct {
		Base
		UserID  uint
		User    User `gorm:"foreignKey:UserID"`
		BlogID  uint
		Blog    Blog   `gorm:"foreignKey:BlogID"`
		Content string `gorm:"not null;type:varchar(512)"`
		BaseTime
	}

	Like struct {
		Base
		UserID uint
		User   User `gorm:"foreignKey:UserID"`
		BlogID uint
		Blog   Blog `gorm:"foreignKey:BlogID"`
		BaseTime
	}
)

type (
	BlogCache struct {
		CountRemake     int
		CountLike       int
		CountCollecting int
		CountRead       int
	}
)

func BlogMigrate() error {
	MigrateList := make([]interface{}, 0)
	MigrateList = append(MigrateList, &Blog{}, &Remake{}, &Like{})
	if err := global.DB.AutoMigrate(MigrateList...); err != nil {
		return err
	}
	return nil
}

//func TagsMigrate() error {
//	if err := global.DB.AutoMigrate(&Tag{}); err != nil {
//		return err
//	}
//	return nil
//}
//
//func RemakesMigrate() error {
//	if err := global.DB.AutoMigrate(&Remake{}); err != nil {
//		return err
//	}
//	return nil
//}
//
//func TagsMigrate() error {
//	if err := global.DB.AutoMigrate(&Tag{}); err != nil {
//		return err
//	}
//	return nil
//}
//
//func TagsMigrate() error {
//	if err := global.DB.AutoMigrate(&Like{}); err != nil {
//		return err
//	}
//	return nil
//}
