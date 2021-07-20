package Dto

import (
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/Util"
)

type TodoDto struct {
	Id          uint   `json:"todo_id"`
	UserId      uint   `json:"user_id"`
	TodoGroupId uint   `json:"todo_group_id"`
	TodoTitle   string `json:"todo_title"`
	TodoContent string `json:"todo_content"`
	IsFinished  bool   `json:"is_finished"`
	CreatedAt   string `json:"created_at"`
}

func TodTodoDto(todo Model.Todo) TodoDto {
	return TodoDto{
		Id:          todo.TodoId,
		UserId:      todo.UserId,
		TodoTitle:   todo.TodoName,
		TodoContent: todo.TodoContent,
		TodoGroupId: todo.GroupId,
		IsFinished:  todo.IsFinished,
		CreatedAt:   todo.CreatedAt.Format(Util.TimeFormat),
	}
}
func TosdTodoDto(todos []Model.Todo) []TodoDto {
	ts := make([]TodoDto, len(todos))
	for i, v := range todos {
		temp := TodoDto{
			Id:          v.TodoId,
			UserId:      v.UserId,
			TodoTitle:   v.TodoName,
			TodoContent: v.TodoContent,
			TodoGroupId: v.GroupId,
			IsFinished:  v.IsFinished,
			CreatedAt:   v.CreatedAt.Format(Util.TimeFormat),
		}
		ts[i] = temp
	}
	return ts
}
