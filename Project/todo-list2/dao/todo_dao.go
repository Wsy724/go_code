package dao

import (
	"todo-list2/config"
	"todo-list2/model"
)

//封装纯数据库操作（不包含业务逻辑），供 service 调用：

// CreateTodo 创建待办
func CreateTodo(todo *model.Todo) error {
	return config.DB.Create(todo).Error
}

// GetTodoByID 根据ID查询待办
func GetTodoByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := config.DB.Where("id=?", id).First(&todo).Error
	return &todo, err
}

// GetAllTodos 查询所有待办
func GetAllTodos() ([]model.Todo, error) {
	var todos []model.Todo
	err := config.DB.Find(&todos).Error
	return todos, err
}

// UpdateTodo 更新待办（支持标题，内容，状态）
func UpdateTodo(todo *model.Todo) error {
	return config.DB.Save(todo).Error
}

// DeleteTodo 删除待办（软删除）
func DeleteTodo(id uint) error {
	return config.DB.Delete(&model.Todo{}, id).Error
}
