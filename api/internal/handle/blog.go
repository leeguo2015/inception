package handle

import "github.com/gin-gonic/gin"

func BlogAdd(c *gin.Context) {


}

func BlogDelete(c *gin.Context) {

	userID := c.GetUint("userID")
	blogId, err := strconv.Atoi(c.Param("blogID"))
	if err != nil {
		global.Log.Error("参数错误：", err.Error())
		response.Fail(c)
		return
	}
	if err := blog.Delete(uint(blogId), userID); err != nil {
		global.Log.Errorf("博客删除失败 %d: %s,", blogId, err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}


func BlogUpdate(c *gin.Context) {
	newBlog := model.NewBlog()
	newBlog.Title = c.PostForm("title")
	blogID, err := strconv.Atoi(c.Param("blogID"))
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
