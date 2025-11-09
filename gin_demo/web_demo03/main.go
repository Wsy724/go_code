package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {

		//方法一：使用map
		//data := map[string]interface{}{
		//	"name":    "八戒",
		//	"age":     30,
		//	"gender":  "男",
		//	"message": "我要回高老庄！",
		//}
		data := gin.H{
			"name":    "唐僧",
			"age":     30,
			"message": "唐僧洗发爱飘柔",
		}
		c.JSON(http.StatusOK, data)
	})

	//方法二：使用结构体
	r.GET("/json02", func(c *gin.Context) {
		type msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		date := msg{
			"孙悟空",
			"吃俺老孙一棒！",
			20,
		}

		c.JSON(http.StatusOK, date)

	})

	r.Run(":9090")
}
