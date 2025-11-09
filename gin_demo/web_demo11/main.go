package main

//中间件
import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index....in")
	name, ok := c.Get("name") //跨中间件存取值
	if !ok {
		name = "匿名用户"
	}
	c.JSON(http.StatusOK, gin.H{
		"meeage": name,
	})
}

// 定义一个中间件m1 统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1....in")
	//计时
	star := time.Now()
	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(star)
	fmt.Printf("花费多久:%v\n", cost)
	fmt.Println("m1....out")
}

func m2(c *gin.Context) {
	fmt.Println("m2....in")
	c.Set("name", "刘华强") //set函数给上下文c设置一个值
	c.Next()             // 调用后续的处理函数
	fmt.Println("m2....out")
}

// 一般中间件都会使用闭包的方式去写
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//准备其他工作...
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登录的判断
			//if是登录用户
			c.Next()
			//else
			//c.Abort
		} else {
			c.Next()
		}
	}

}

func main() {
	r := gin.Default()

	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件m1,m2,authMiddleware

	//GET(relativePath string, handlers ...HandlerFunc)

	//r.GET("/index", m1, indexHandler)
	r.GET("/index", indexHandler)

	r.GET("/page", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"meeage": "page",
		})
	})

	//路由组注册中间件方法1：
	//xx1 := r.Group("/xx1", authMiddleware(false))
	//{
	//	xx1.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{"meeage": "xx1/index"})
	//	})
	//}
	//
	////路由组注册中间件方法2：
	//xx2 := r.Group("/xx2")
	//xx2.Use(authMiddleware(false))
	//{
	//	xx2.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{"meeage": "xx2/index"})
	//	})
	//}

	r.Run(":8080")
}
