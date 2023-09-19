package handle

import (
	"inception/api/internal/global"
	"inception/api/internal/logic/blog"
	"inception/api/internal/response"
	"inception/api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentAdd(c *gin.Context) {
	commentType := blog.CommentType(c.PostForm("comment_type"))
	content := c.PostForm("content")
	userID, _ := strconv.Atoi(c.PostForm("user_id"))
	ownerID, _ := strconv.Atoi(c.PostForm("owner_id"))
	err := blog.CommentAdd(commentType, content, uint(userID), uint(ownerID))
	if err != nil {
		global.Log.Error(err)
		response.Fail(c)
	}
	response.Ok(c)

}

func CommentDelete(c *gin.Context) {
	userID := c.GetUint(utils.UserID)
	commentID, _ := strconv.Atoi(c.PostForm("omment_id"))
	if err := blog.CommentDelete(uint(commentID), userID); err != nil {
		global.Log.Error(err)
		response.Fail(c)
	}
	response.Ok(c)
}

func CommentGet(c *gin.Context) {
	blogID, _ := strconv.Atoi(c.PostForm("blog_id"))
	start, _ := strconv.Atoi(c.Query("start"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	comments, count, err := blog.CommentGet(uint(blogID), start, limit)
	if err != nil {
		global.Log.Error(err)
		response.Fail(c)
	}

	data := map[string]interface{}{
		"data":  comments,
		"count": count,
	}
	response.OkWithData(data, c)

}
