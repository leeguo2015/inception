package blog

import (
	"fmt"
	. "inception/api/internal/global"
	"inception/api/internal/model"
)

func CommentAdd(commentType model.CommentType, content string, blogID, userID, ownerID uint) error {
	comment := new(model.Comment)
	comment.BlogID = blogID
	comment.Content = content
	comment.UserID = userID
	comment.CommentType = commentType
	switch commentType {
	case model.CommentTypeBlog:
		// comment.BlogID = ownerID
	case model.CommentTypeComment:
		comment.ParentID = ownerID
	default:
		return fmt.Errorf("unsupported comment type %s", commentType)
	}

	if err := DB.Create(comment).Error; err != nil {
		return err
	}
	if commentType == model.CommentTypeComment {
		var parent model.Comment
		DB.Model(&parent).Where("ID = ?", ownerID).First(&parent)
		parent.Replies = append(parent.Replies, comment.ID)
		DB.Save(&parent)
	}
	return nil
}

func CommentDelete(CommentID uint, userID uint) error {
	comment := new(model.Comment)
	err := DB.Where("id = ? AND user_id = ?", CommentID, userID).Delete(comment).Error
	if err != nil {
		return fmt.Errorf("delete comment failed %s", err)
	}
	return nil

}

func CommentGet(blogID uint, start, limit int) ([]model.Comment, int64, error) {
	comments := []model.Comment{}
	total := int64(0)

	// 查询符合条件的评论
	err := DB.Where("blog_id = ?", blogID).Offset(start).Limit(limit).Find(&comments).Error

	if err != nil {
		return nil, total, err
	}

	// 查询总评论数
	err = DB.Model(&model.Comment{}).Where("blog_id = ?", blogID).Count(&total).Error

	if err != nil {
		return nil, total, err
	}

	return comments, total, nil
}
