package handle

import (
	"inception/api/internal/model"
	"inception/api/internal/response"

	"github.com/gin-gonic/gin"
)

func RemarkAdd(c *gin.Context) {

	response.Ok(c)

}

func RemarkDelete(c *gin.Context) {

	response.Ok(c)
}

func RemarkUpdate(c *gin.Context) {

	response.Ok(c)
}

func RemarkGet(c *gin.Context) {
	remake := model.Remake{}

	response.OkWithData(remake, c)

}
