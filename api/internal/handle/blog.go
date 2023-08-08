package handle

import (
	"inception/api/internal/global"
	"inception/api/internal/logic/blog"
	"inception/api/internal/model"
	"inception/api/internal/response"
	"inception/api/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogAdd(c *gin.Context) {
	claims, err := utils.GetClaims(c)
	if err != nil {
		global.Log.Error("获取用户信息错误：", err.Error())
		return
	}
	newBlog := model.NewBlog()

	newBlog.Title = c.PostForm("title")
	newBlog.Content = c.PostForm("content")
	tags := blog.GetInsertTags(c.PostFormArray("tags"))
	// global.Log.Info("c.PostFormArray", tags[0].ID, tags[1].ID)

	//newBlog.Tags = blog.GetInsertTags(c.PostFormArray("tags"))
	newBlog.UserID = claims.ID
	global.Log.Info(newBlog, claims)
	if err := blog.Add(claims.ID, newBlog, tags); err != nil {
		global.Log.Errorf("%s 博客添加失败:%s,", claims.Username, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

func BlogDelete(c *gin.Context) {

	blogId, err := strconv.Atoi(c.Param("blogID"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	if err := blog.Delete(uint(blogId)); err != nil {
		global.Log.Errorf("博客删除失败 %d: %s,", blogId, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func BlogUpdate(c *gin.Context) {

}

func BlogGet(c *gin.Context) {

	blogId, err := strconv.Atoi(c.Param("blogID"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	blog, err := blog.Get(uint(blogId))
	if err != nil {
		global.Log.Errorf("博客删除失败blogId: %d: %s,", blogId, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blog, c)

}
