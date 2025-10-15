package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"venturan/global"
)

// 响应结构体
type Response struct {
	Code int         `json:"code"` // 业务码
	Data interface{} `json:"data"` // 数据
	Msg  string      `json:"msg"`  // 提示消息
}

// Success 响应成功 code 为 200 表示成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		200,
		data,
		"",
	})
}

// Fail 响应失败 code 不为 200 表示失败
func Fail(c *gin.Context, errorCode int, msg string) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		nil,
		msg,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}
