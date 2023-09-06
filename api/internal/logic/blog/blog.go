package blog

import (
	"errors"
	"fmt"
	"inception/api/internal/global"
	"inception/api/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
//	@Description: 不在该层做鉴权 1.删除标签，2，删除点赞评论，3，删除博客。
//	@param blogID
//	@return error
func Delete(blogID uint, UserID uint) error {
	blog := new(model.Blog)
	blog.ID = blogID
	blog.UserID = UserID
	// if err := global.DB.Where("id=", blogID).First(blog).Error; err != nil {
	// 	return fmt.Errorf("get blog failed %s", err)
	// }
	// global.DB.Where("id=", blogID).Delete(&model.Like{})
	// global.DB.Where("id=", blogID).Delete(&model.Comment{})
	// err := global.DB.Delete(blog).Error
	err := global.DB.Select(clause.Associations).Delete(blog).Error
	if err != nil {
		return fmt.Errorf("delete blog failed %s", err)
	}
	return nil
}

//方案。 1.直接更新所有内容， 2.局部更新，（1）更新标签，（2）内容/标题/

func Update(UserId uint, blog *model.Blog, tags []*model.Tag) error {

	blog.UserID = UserId
	if err := global.DB.Model(&blog).Updates(map[string]interface{}{"title": blog.Title, "content": blog.Content}).Error; err != nil {
		return err
	}
	if err := global.DB.Model(&blog).Association("Tags").Replace(tags); err != nil {
		return err
	}
	return nil
}

func Get(UserId uint) (model.Blog, error) {
	blog := model.NewBlog()
	if global.DB.Where("id = ?", UserId).First(blog).Error == gorm.ErrRecordNotFound {
		return model.Blog{}, errors.New("博客内容不存在")
	}
	return *blog, nil
}

