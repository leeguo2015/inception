package blog

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"inception/api/internal/global"
	"inception/api/internal/model"
)

func Add(UserId uint, blog *model.Blog, tags []*model.Tag) error {
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

// Delete
//
//	@Description:
//
// 1.删除标签，2，删除点赞评论，3，删除博客。
//
//	@param blogID
//	@return error
func Delete(blogID uint) error {
	blog := new(model.Blog)
	if err := global.DB.Where("id=", blogID).First(blog).Error; err != nil {
		return fmt.Errorf("get blog failed %s", err)
	}
	global.DB.Where("id=", blogID).Delete(&model.Like{})
	global.DB.Where("id=", blogID).Delete(&model.Remake{})
	global.DB.Delete(&model.Remake{})
	err := global.DB.Delete(blog).Error
	if err != nil {
		return fmt.Errorf("delete blog failed %s", err)
	}
	return nil
}
