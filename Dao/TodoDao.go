package Dao

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
)

func GetTodoByTodoName(todoName string) (*Model.Todo, error) {
	todo := new(Model.Todo)
	if err := Db.DB.Where("todo_name = ?", todoName).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoById(todoId uint) (*Model.Todo, error) {
	todo := new(Model.Todo)
	if err := Db.DB.Where("todo_id=?", todoId).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoByGroupID(id uint) ([]*Model.Todo, error) {
	var todos []*Model.Todo
	err := Db.DB.Where("group_id=?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}
func GetTodoByUserID(id uint) ([]*Model.Todo, error) {
	var todos []*Model.Todo
	err := Db.DB.Where("UserId=?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func AddTodo(todo *Model.Todo) (*Model.Todo, error) {
	err := Db.DB.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func UpdateTodo(todo *Model.Todo) (*Model.Todo, error) {
	err := Db.DB.Omit("created_at", "user_id", "group_id", "todo_name", "todo_content").Save(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func DeleteTodoById(id, userId uint) error {
	todo := Model.Todo{}
	err := Db.DB.Where("todo_id = ? and user_id = ?", id, userId).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteTodo(todo *Model.Todo) (*Model.Todo, error) {
	err := Db.DB.Delete(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func DeleteTodoByName(name string) error {
	todo := Model.Todo{}
	err := Db.DB.Where("todo_name = ?", name).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
