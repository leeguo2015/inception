package blog

import (
	"errors"
	"gorm.io/gorm"
	"inception/api/internal/global"
	"inception/api/internal/model"
)

func Add(UserId uint, blog *model.Blog, tags []model.Tag) error {
	user := new(model.User)
	if global.DB.Where("id = ?", UserId).First(user).Error == gorm.ErrRecordNotFound {
		return errors.New("用户不存在")
	}
	blog.UserID = user.ID
	if err := global.DB.Create(&blog).Error; err != nil {
		return err
	}
	if err := global.DB.Model(&blog).Association("Tags").Append(tags); err != nil {
		return err
	}
	return nil
}

func Delete(blogID uint) error {
	blog := new(model.Blog)
	err := global.DB.Delete("id = ?", blogID).First(blog).Error
	if err != nil {
		global.Log.Error("删除博客失败", err)
		return errors.New("删除博客失败")
	}
	return nil
}
