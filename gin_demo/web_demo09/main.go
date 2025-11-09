package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "ok",
		//})

		////状态码 http.StatusMovedPermanently,值为 301,表示永久重定向
		//跳转到github网站
		c.Redirect(http.StatusMovedPermanently, "https://github.com/")
	})

	r.GET("/a", func(c *gin.Context) {
		//跳转到b路由处理函数
		c.Request.URL.Path = "/b" //把请求的URL修改
		//继续后续的处理
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "已经跳转到B执行函数内",
		})
	})

	r.Run(":8080")
}
