package handle

import (
	. "inception/api/internal/global"
	"inception/api/internal/response"
	"inception/api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Uploads(c *gin.Context) {
	file, err := c.FormFile("file")
	Log.Debug(file)
	if err != nil {
		Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 保存文件到本地
	err = c.SaveUploadedFile(file, utils.StaticURL+file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	response.Ok(c)

}

