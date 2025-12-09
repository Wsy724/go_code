package router

import (
	"todo-list2/controller"

	"github.com/gin-gonic/gin"
)

//绑定接口和处理函数，定义 API 路径：

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default() //默认包含日志，恢复中间件

	//待办事项接口分组
	todoGroup := r.Group("/api/todo")
	{
		todoGroup.POST("/create", controller.CreateTodo)       //创建待办
		todoGroup.GET("/get/:id", controller.GetTodoByID)      //根据ID查询
		todoGroup.GET("/list", controller.GetAllTodos)         //查询所有
		todoGroup.PUT("/update/:id", controller.UpdateTodo)    //更新待办
		todoGroup.DELETE("/delete/:id", controller.DeleteTodo) //删除待办
	}
	return r
}
