package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//请求方法大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"GET": "ok"})

		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"POST": "ok"})

		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"PUT": "ok"})
		}

	})

	//NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "没有这个路由~"})
	})

	//路由组
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"shop/index": "ok"})
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"shop/cart": "ok"})
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"shop/checkout": "ok"})
		})

		//嵌套路由组
		xx := shopGroup.Group("/xx")
		{
			xx.GET("/index", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"shop/xx/index": "ok"})
			})
		}

	}

	r.Run(":9090")
}
