// Package response
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-16 17:47
package response

import (
	"cnic/fairman-gin/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// start time
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func ResultAll(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(code, Response{
		code,
		data,
		message,
	})
}

func Ok(c *gin.Context) {
	Result(errno.SUCCESS, map[string]interface{}{}, "Operation successful", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(errno.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(errno.SUCCESS, data, "Operation successful", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errno.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(errno.ERROR, map[string]interface{}{}, "operation failed", c)
}

func FailCode(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "operation failed", c)
}

func FailCodeMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(errno.ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errno.ERROR, data, message, c)
}

func FailTokenWithMessage(message string, c *gin.Context) {
	Result(errno.TIMEOUT, map[string]interface{}{}, message, c)
}
