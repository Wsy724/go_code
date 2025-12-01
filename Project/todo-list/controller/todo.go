package controller

import (
	"strconv"
	"todo-list/service"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

// CreateTodo 创建待办
func CreateTodo(c *gin.Context) {
	// 解析请求参数
	var req struct {
		Title   string `json:"title" binding:"required"` // 必传
		Content string `json:"content"`
	}
	// 参数绑定失败（比如title为空）
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误："+err.Error())
		return
	}
	// 调用业务逻辑
	err := service.CreateTodo(req.Title, req.Content)
	if err != nil {
		utils.Fail(c, "创建失败："+err.Error())
		return
	}
	utils.Success(c, nil)
}

// GetTodoByID 根据ID查询待办
func GetTodoByID(c *gin.Context) {
	// 从URL获取id参数
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ParamError(c, "ID格式错误")
		return
	}
	// 调用业务逻辑
	todo, err := service.GetTodoByID(uint(id))
	if err != nil {
		utils.Fail(c, "查询失败："+err.Error())
		return
	}
	utils.Success(c, todo)
}

// GetAllTodos 查询所有待办
func GetAllTodos(c *gin.Context) {
	todos, err := service.GetAllTodos()
	if err != nil {
		utils.Fail(c, "查询失败："+err.Error())
		return
	}
	utils.Success(c, todos)
}

// UpdateTodo 更新待办
func UpdateTodo(c *gin.Context) {
	// 解析URL中的ID
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ParamError(c, "ID格式错误")
		return
	}
	// 解析请求参数
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Status  int    `json:"status" binding:"oneof=0 1"` // 只能是0或1
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误："+err.Error())
		return
	}
	// 调用业务逻辑
	err = service.UpdateTodo(uint(id), req.Title, req.Content, req.Status)
	if err != nil {
		utils.Fail(c, "更新失败："+err.Error())
		return
	}
	utils.Success(c, nil)
}

// DeleteTodo 删除待办
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ParamError(c, "ID格式错误")
		return
	}
	err = service.DeleteTodo(uint(id))
	if err != nil {
		utils.Fail(c, "删除失败："+err.Error())
		return
	}
	utils.Success(c, nil)
}
