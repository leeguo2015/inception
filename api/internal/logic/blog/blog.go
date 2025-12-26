package blog

import (
	"errors"
	"fmt"
	"inception/api/internal/global"
	"inception/api/internal/model"
	"strconv"
	"strings"

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

func GetList(page, pageSize int, categoryIds string) ([]model.Blog, int64, error) {
	var blogs []model.Blog
	var total int64
	
	// 计算偏移量
	offset := (page - 1) * pageSize
	
	// 构建查询条件
	db := global.DB.Model(&model.Blog{}).Preload("Tags").Preload("User")
	
	// 如果有分类筛选
	if categoryIds != "" {
		// 这里需要根据categoryIds来筛选关联的博客
		// 由于是多对多关系，需要先查询属于这些分类的博客ID
		var blogIDs []uint
		if err := global.DB.Table("blog_tags").
			Where("tag_id IN ?", parseCategoryIds(categoryIds)).
			Pluck("blog_id", &blogIDs).Error; err != nil {
			return nil, 0, err
		}
		if len(blogIDs) > 0 {
			db = db.Where("id IN ?", blogIDs)
		}
	}
	
	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 获取分页数据
	if err := db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, 0, err
	}
	
	return blogs, total, nil
}

// parseCategoryIds 解析分类ID字符串，支持多种格式
func parseCategoryIds(categoryIds string) []uint {
	// 这里可以根据实际前端传递的格式进行解析
	// 比如 "[1,2,3]" 或 "1,2,3" 等格式
	// 简化处理：假设是逗号分隔的字符串
	var ids []uint
	if categoryIds != "" {
		// 移除可能的方括号
		cleanIds := strings.Trim(categoryIds, "[]")
		// 按逗号分割
		idStrs := strings.Split(cleanIds, ",")
		for _, idStr := range idStrs {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32); err == nil {
				ids = append(ids, uint(id))
			}
		}
	}
	return ids
}

