package handle

import (
	"inception/api/internal/model"
	"inception/api/internal/response"

	"github.com/gin-gonic/gin"
)

func CommentAdd(c *gin.Context) {

	response.Ok(c)

}

func CommentDelete(c *gin.Context) {

	response.Ok(c)
}

// func CommentUpdate(c *gin.Context) {

// 	response.Ok(c)
// }

func CommentGet(c *gin.Context) {
	Comment := model.Comment{}
	
	response.OkWithData(Comment, c)

}
