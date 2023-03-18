package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 400
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func httpResult(httpStatus int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpStatus, gin.H{
		"data": data,
		"msg":  msg,
	})
}

//使用http原生协议

func HttpOk(msg string, c *gin.Context) {
	httpResult(http.StatusOK, nil, msg, c)
}

func HttpOkDetails(data interface{}, msg string, c *gin.Context) {
	httpResult(http.StatusOK, data, msg, c)
}

func HttpBadDetail(data interface{}, msg string, c *gin.Context) {
	httpResult(http.StatusBadRequest, data, msg, c)
}
func HttpBadRequest(msg string, c *gin.Context) {
	httpResult(http.StatusBadRequest, nil, msg, c)
}

func HttpUnauthorized(data interface{}, msg string, c *gin.Context) {
	httpResult(http.StatusUnauthorized, data, msg, c)
}

func HttpNotFound(data interface{}, msg string, c *gin.Context) {
	httpResult(http.StatusNotFound, data, msg, c)
}

func HttpUnprocessableEntity(msg string, c *gin.Context) {
	httpResult(http.StatusUnprocessableEntity, nil, msg, c)
}
