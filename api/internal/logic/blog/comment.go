package blog

import "inception/api/internal/model"

type CommentType string

var CommentTypeBlog CommentType = "blog"
var CommentTypeComment CommentType = "comment"

func CommentAdd(commentType CommentType, comment string， UserID, uint) error {
	return nil
}

func CommentDelete() {

}

func CommentGet(blogID uint, start, limit int) (model.Comment, int, error) {
	Comment := model.Comment{}

	return Comment, 0, nil
}
