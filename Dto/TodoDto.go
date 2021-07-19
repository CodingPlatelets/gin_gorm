package Dto

import (
	"github.com/WenkanHuang/gin_gorm/Model"
	"time"
)

type TodoDto struct {
	Id          uint      `json:"todo_id"`
	UserId      uint      `json:"user_id"`
	TodoGroupId uint      `json:"todo_group_id"`
	TodoTitle   string    `json:"todo_title"`
	TodoContent string    `json:"todo_content"`
	IsFinished  bool      `json:"is_finished"`
	CreatedAt   time.Time `json:"created_at"`
}

func TodTodoDto(todo Model.Todo) TodoDto {
	return TodoDto{
		Id:          todo.TodoId,
		UserId:      todo.UserId,
		TodoTitle:   todo.TodoName,
		TodoContent: todo.TodoContent,
		TodoGroupId: todo.GroupId,
		IsFinished:  todo.IsFinished,
		CreatedAt:   todo.CreatedAt,
	}
}
