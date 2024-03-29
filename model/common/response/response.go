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
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	HttpResponse(http.StatusOK, code, data, msg, c)
}

func HttpResponse(status, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(status, Response{
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

func FailAuthExpireWithDetailed(data interface{}, message string, c *gin.Context) {
	HttpResponse(http.StatusUnauthorized, http.StatusUnauthorized, data, message, c)
}

func FailStatusForbiddenWithDetailed(data interface{}, message string, c *gin.Context) {
	HttpResponse(http.StatusForbidden, http.StatusForbidden, data, message, c)
}

func FailStatusTooManyRequestsWithDetailed(data interface{}, message string, c *gin.Context) {
	HttpResponse(http.StatusTooManyRequests, http.StatusTooManyRequests, data, message, c)
}
