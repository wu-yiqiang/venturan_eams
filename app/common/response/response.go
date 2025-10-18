package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
func Fail(c *gin.Context, errorType Response) {
	c.JSON(http.StatusOK, Response{
		errorType.Code,
		errorType.Data,
		errorType.Msg,
	})
}
