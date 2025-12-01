package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应结构体
type Response struct {
	Code int         `json:"code"` // 状态码：200成功，500失败，400参数错误
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 响应数据（可选）
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 500,
		Msg:  msg,
		Data: nil,
	})
}

// ParamError 参数错误响应
func ParamError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}
