package Controller

import (
	"github.com/WenkanHuang/gin_gorm/Dao"
	"github.com/WenkanHuang/gin_gorm/Dto"
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/Response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddTodo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(Model.User).UserId
	var todo Model.Todo
	err := ctx.ShouldBind(&todo)
	if err != nil {
		Response.Fail(ctx, gin.H{"error": err}, "bind failed")
		return
	} else {
		todo.UserId = userId
	}
	_, errAdd := Dao.AddTodo(&todo)
	if errAdd != nil {
		Response.Fail(ctx, gin.H{"error": errAdd}, "add failed")
		return
	}
	count, errUser := Dao.GetGroupByGroupId(todo.GroupId)
	if errUser != nil {
		Response.Fail(ctx, gin.H{"error": errUser}, "group is not exits")
	} else {
		count.ItemCOUNT++
		_, errCount := Dao.UpdateGroup(count)
		if errCount != nil {
			Response.Fail(ctx, gin.H{"error": errCount}, "add item failed")
		}
	}
	Response.Success(ctx, gin.H{"Todo": Dto.TodTodoDto(todo)}, "OK")

}

func UpdateTodo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	id, errFormat := strconv.Atoi(ctx.Param("id"))
	if errFormat != nil {
		Response.Fail(ctx, gin.H{"error": errFormat}, "id is not a integer value")
	} else {
		var bind Model.Todo
		errBind := ctx.ShouldBind(&bind)
		if errBind != nil {
			Response.Fail(ctx, gin.H{"error": errBind.Error()}, "bind failed")
			return
		} else {
			bind.TodoId = uint(id)
			bind.UserId = user.(Model.User).UserId
			_, errUpdate := Dao.UpdateTodo(&bind)
			if errUpdate != nil {
				Response.Fail(ctx, gin.H{"error": errUpdate}, "update failed")
				return
			} else {
				Response.Success(ctx, nil, "OK")
			}

		}
	}

}

func DeleteTodo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(Model.User).UserId
	todoId, errFormat := strconv.Atoi(ctx.Param("id"))
	if errFormat != nil {
		Response.Fail(ctx, gin.H{"error": errFormat}, "id is not an integer value")
	} else {
		count, errGroup := Dao.GetGroupByTodoId(uint(todoId))
		if errGroup != nil {
			Response.Fail(ctx, gin.H{"error": errGroup}, "group is not exits")
		} else {
			count.ItemCOUNT--
			_, errCount := Dao.UpdateGroup(count)
			if errCount != nil {
				Response.Fail(ctx, gin.H{"error": errCount}, "delete item failed")
			}
		}
		errDelete := Dao.DeleteTodoById(uint(todoId), userId)
		if errDelete != nil {
			Response.Fail(ctx, gin.H{"error": errDelete}, "delete failed")
		} else {
			Response.Success(ctx, nil, "OK")
		}
	}
}
