package dao

import (
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/dto"
	"github.com/WenkanHuang/gin_gorm/model"
)

func GetTodoByTodoName(todoName string) (*model.Todo, error) {
	todo := new(model.Todo)
	if err := db.DB.Where("todo_name = ?", todoName).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoById(todoId uint) (*model.Todo, error) {
	todo := new(model.Todo)
	if err := db.DB.Where("todo_id=?", todoId).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoByGroupID(id uint) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := db.DB.Where("group_id=?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func GetTodoByUserID(id uint) ([]*model.Todo, error) {
	var todos []*model.Todo
	err := db.DB.Where("UserId=?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func AddTodo(todo *model.Todo) error {
	err := db.DB.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	if todo.IsFinished == false {
		err := db.DB.Model(todo).Update("is_finished", 0).Error
		if err != nil {
			return nil, err
		}
	}
	err2 := db.DB.Omit("created_at" /* "user_id", "group_id", "todo_name", "todo_content"*/).Updates(todo).Error
	if err2 != nil {
		return nil, err2
	}
	return todo, nil
}

func DeleteTodoById(id, userId uint) error {
	todo := model.Todo{}
	err := db.DB.Where("todo_id = ? and user_id = ?", id, userId).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *model.Todo) (*model.Todo, error) {
	err := db.DB.Delete(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func DeleteTodoByName(name string) error {
	todo := model.Todo{}
	err := db.DB.Where("todo_name = ?", name).Delete(&todo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTodoBySelectCondition(c dto.Condition) ([]model.Todo, error) {
	var todo []model.Todo
	sql := db.DB.Where("created_at < ?", c.CreatedAt)

	// CAUTION may have injection
	sel := "%" + c.Keyword + "%"
	sql = db.DB.Model(&todo).Where("todo_content like ?", sel)
	if c.GroupId != 0 {
		sql = db.DB.Model(&todo).Where("group_id = ? ", c.GroupId)
	}
	if c.IsFinished == true {
		sql = db.DB.Model(&todo).Where("is_finished = ?", 1)
	} else {
		sql = db.DB.Model(&todo).Where("is_finished = ?", 0)
	}
	sql.Find(&todo)
	return todo, nil
}
