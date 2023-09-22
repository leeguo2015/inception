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
	newBlog.UserID = claims.UserID
	global.Log.Info(newBlog, claims)
	if err := blog.Add(claims.UserID, newBlog, tags); err != nil {
		global.Log.Errorf("%s 博客添加失败:%s,", claims.Username, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func BlogDelete(c *gin.Context) {
	userID := c.GetUint("userID")
	blogID, err := strconv.Atoi(c.Param("blog_id"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	if err := blog.Delete(uint(blogID), userID); err != nil {
		global.Log.Errorf("博客删除失败 %d: %s,", blogID, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func BlogUpdate(c *gin.Context) {
	newBlog := model.NewBlog()
	newBlog.Title = c.PostForm("title")
	blogID, err := strconv.Atoi(c.Param("blog_id"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	newBlog.ID = uint(blogID)
	newBlog.Content = c.PostForm("content")
	tags := blog.GetInsertTags(c.PostFormArray("tags"))
	if err := blog.Update(c.GetUint(utils.UserID), newBlog, tags); err != nil {
		global.Log.Errorf("userID: %d 更新博客失败:%s,", c.GetUint(utils.UserID), err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func BlogGet(c *gin.Context) {
	blogID, err := strconv.Atoi(c.Param("blog_id"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	blog, err := blog.Get(uint(blogID))
	if err != nil {
		global.Log.Errorf("博客删除失败blogID: %d: %s,", blogID, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blog, c)

}
