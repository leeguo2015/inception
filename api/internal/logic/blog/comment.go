package blog

import (
	"fmt"
	"inception/api/internal/global"
	"inception/api/internal/model"

	"gorm.io/gorm/clause"
)

type CommentType string

var CommentTypeBlog CommentType = "blog"
var CommentTypeComment CommentType = "comment"

func CommentAdd(commentType CommentType, Comtent string, UserID, OwnerID uint) error {
	comment := new(model.Comment)
	switch commentType {
	case CommentTypeBlog:
		comment.BlogID = OwnerID
	case CommentTypeComment:
		comment.ParentID = OwnerID
	default:
		return fmt.Errorf("unsupported comment type %s", commentType)
	}
	comment.Content = Comtent
	comment.UserID = UserID

	if err := global.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func CommentDelete(CommentID uint) error {
	comment := new(model.Comment)
	comment.ID = CommentID
	err := global.DB.Select(clause.Associations).Delete(comment).Error
	if err != nil {
		return fmt.Errorf("delete comment failed %s", err)
	}
	return nil

}


func CommentGet(blogID uint, start, limit int) ([]model.Comment, int64, error) {
	comments := []model.Comment{}
	total := int64(0)

	// 查询符合条件的评论
	err := global.DB.Where("blog_id = ?", blogID).Offset(start).Limit(limit).Find(&comments).Error

	if err != nil {
		return nil, total, err
	}

	// 查询总评论数

	err = global.DB.Model(&model.Comment{}).Where("blog_id = ?", blogID).Count(&total).Error

	if err != nil {
		return nil, total, err
	}

	return comments, total, nil
}
