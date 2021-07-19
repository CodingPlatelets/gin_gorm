package Dao

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"
)

func GetTodoByTodoName(todoName string) (*Model.Todo, error) {
	todo := new(Model.Todo)
	if err := Db.DB.Where("todoName=?", todoName).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoById(todoId uint) (*Model.Todo, error) {
	todo := new(Model.Todo)
	if err := Db.DB.Where("todoId=?", todoId).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoByGroupID(id uint) ([]*Model.Todo, error) {
	var todos []*Model.Todo
	err := Db.DB.Where("GroupId=?", id).Find(&todos).Error
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
	err := Db.DB.Save(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func DeleteTodoById(id uint) error {
	todo := Model.Todo{}
	err := Db.DB.Where("todoId = ?", id).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteTodoByName(name string) error {
	todo := Model.Todo{}
	err := Db.DB.Where("todoName = ?", name).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}
