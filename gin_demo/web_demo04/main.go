package main

//querystring

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//GET求取 URL ?后面的是querystring参数
	//key=value格式，多个key-value用 &连接
	//er:  /web/query=杨超越&age18
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发送的query string参数
		name := c.Query("query") //通过Query获取 query string的参数
		age := c.Query("age")
		//name := c.DefaultQuery("query", "somebody")//获取不到就用指定的默认值
		//name, ok := c.GetQuery("query") //取到值返回(值,ture)  ,取不到返回("",false)
		//if !ok {
		//	//取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}
