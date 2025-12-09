package service

//处理业务逻辑（参数校验、调用 DAO），隔离路由和数据库操作：
import (
	"fmt"
	"todo-list2/dao"
	"todo-list2/model"
)

// CreateTodo 创建待办的业务逻辑
func CreateTodo(title, content string) error {
	//简单参数校验
	if title == "" {
		return fmt.Errorf("待办标题不能为空")
	}
	todo := &model.Todo{
		Title:   title,
		Content: content,
	}
	return dao.CreateTodo(todo)
}

// GetTodoByID 根据ID查询待办
func GetTodoByID(id uint) (*model.Todo, error) {
	return dao.GetTodoByID(id)
}

// GetAllTodo 查询所有待办
func GetAllTodos() ([]model.Todo, error) {
	return dao.GetAllTodos()
}

// UpdateTodo 更新待办
func UpdateTodo(id uint, title, content string, status int) error {
	//先查询待办是否存在
	todo, err := dao.GetTodoByID(id)
	if err != nil {
		return fmt.Errorf("待办不存在:%v", err)
	}

	//更新字段
	if title != "" {
		todo.Title = title
	}
	if content != "" {
		todo.Content = content
	}
	todo.Status = status
	return dao.UpdateTodo(todo)
}

// DeleteTodo 删除待办（软删除）
func DeleteTodo(id uint) error {
	return dao.DeleteTodo(id)
}
