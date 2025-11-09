package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	///login post请求
	r.POST("/login", func(c *gin.Context) {
		//username := c.PostForm("username") //"username" 就是根据 form 表单中 input 元素的 name 属性来确定的。
		//password := c.PostForm("password")
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "SB"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "****"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})

	r.Run(":9090")
}
