package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //创建一个默认的路由

	//加载静态文件
	r.Static("/xxx", "./statics") //

	//gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//r.LoadHTMLFiles("templates/index.tmpl") //模板解析
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模板渲染
			"title": "posts/index.tmpl",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{ //模板渲染
			"title": "<a href=:'https://https://lol.qq.com/main.shtml'>英雄联盟</a>",
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.Run(":9090") //启动sever
}
