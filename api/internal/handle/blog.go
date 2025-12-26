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

type BlogAddRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags"`
}

func BlogAdd(c *gin.Context) {
	claims, err := utils.GetClaims(c)
	if err != nil {
		global.Log.Error("获取用户信息错误：", err.Error())
		response.FailWithMessage("认证失败", c)
		return
	}
	
	var req BlogAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("参数绑定错误：", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}

	newBlog := model.NewBlog()
	newBlog.Title = req.Title
	newBlog.Content = req.Content
	tags := blog.GetInsertTags(req.Tags)
	newBlog.UserID = claims.UserID
	
	global.Log.Info("添加博客:", newBlog, claims)
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

type BlogUpdateRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Tags    []string `json:"tags"`
}

func BlogUpdate(c *gin.Context) {
	claims, err := utils.GetClaims(c)
	if err != nil {
		global.Log.Error("获取用户信息错误：", err.Error())
		response.FailWithMessage("认证失败", c)
		return
	}
	
	var req BlogUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Log.Error("参数绑定错误：", err.Error())
		response.FailWithMessage("参数错误", c)
		return
	}

	blogID, err := strconv.Atoi(c.Param("blog_id"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	
	newBlog := model.NewBlog()
	newBlog.ID = uint(blogID)
	newBlog.Title = req.Title
	newBlog.Content = req.Content
	tags := blog.GetInsertTags(req.Tags)
	
	if err := blog.Update(claims.UserID, newBlog, tags); err != nil {
		global.Log.Errorf("userID: %d 更新博客失败:%s,", claims.UserID, err.Error())
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
		global.Log.Errorf("博客获取失败blogID: %d: %s,", blogID, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(blog, c)
}

type BlogListRequest struct {
	Page       int    `form:"page" binding:"min=1"`
	PageSize   int    `form:"pageSize" binding:"min=1,max=50"`
	CategoryIds string `form:"categoryIds"`
}

func BlogList(c *gin.Context) {
	var req BlogListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		global.Log.Error("博客列表参数错误：", err.Error())
		response.HttpBadRequest("参数格式错误", c)
		return
	}
	
	blogs, total, err := blog.GetList(req.Page, req.PageSize, req.CategoryIds)
	if err != nil {
		global.Log.Error("获取博客列表失败:", err.Error())
		response.FailWithMessage("获取博客列表失败", c)
		return
	}
	
	response.OkWithData(gin.H{
		"list": blogs,
		"total": total,
		"page": req.Page,
		"pageSize": req.PageSize,
	}, c)
}
