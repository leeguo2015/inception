package model

import (
	"inception/api/internal/global"

	"github.com/google/uuid"

)

type (
	Tag struct {
		Base

		Name string `gorm:"not null;type:varchar(48);uniqueIndex"`
		BaseTime
		Blogs []Blog `gorm:"many2many:blog_tags;" json:"-"`

	}

	Blog struct {
		Base

		Uuid   uuid.UUID `json:"uuid" gorm:"index;column:uuid"`
		UserID uint      `gorm:"not null"`
		User   User      `gorm:"foreignKey:UserID" json:"-"`


		Title   string `gorm:"not null;type:varchar(48)"`
		Content string ` gorm:"not null;type:text"`
		BaseTime


		Comment []Comment `json:"-"`
		Likes   []Like    `json:"-"`
		Tags    []Tag     `gorm:"many2many:blog_tags;"`
	}

	Comment struct {
		Base
		UserID   uint
		User     User `json:"-" gorm:"foreignKey:UserID"`
		BlogID   uint
		Blog     Blog      `gorm:"foreignKey:BlogID"`
		Content  string    `gorm:"not null;type:varchar(512)"`
		ParentID uint      // 表示父评论的ID
		Replies  []Comment `gorm:"foreignkey:ParentID"` // 子评论，即多个评论的评论
		BaseTime
	}

	Like struct {
		Base
		UserID uint

		User   User `json:"-" gorm:"foreignKey:UserID"`
		BlogID uint
		Blog   Blog `json:"-" gorm:"foreignKey:BlogID" `

		BaseTime
	}
)

type (
	BlogCache struct {
		CountComment    int
		CountLike       int
		CountCollecting int
		CountRead       int
	}
)


const (
	TagTableName  = "tags"
	BlogTableName = "blogs"
	UserTableName = "users"
)

func NewBlog() *Blog {
	return &Blog{
		Uuid: uuid.New(),
	}
}

func BlogMigrate() error {
	MigrateList := make([]interface{}, 0)
	MigrateList = append(MigrateList, &Blog{}, &Comment{}, &Like{})

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

