package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/WenkanHuang/gin_gorm/dao"
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/dto"
	"github.com/WenkanHuang/gin_gorm/model"
	"github.com/WenkanHuang/gin_gorm/response"
	"github.com/WenkanHuang/gin_gorm/util"
)

func AddTodo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	var todo model.Todo
	err := ctx.ShouldBind(&todo)
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "bind failed")
		return
	}
	todo.UserId = userId

	errAdd := dao.AddTodo(&todo)
	if errAdd != nil {
		response.Fail(ctx, gin.H{"error": errAdd}, "add failed")
		return
	}
	count, errUser := dao.GetGroupByGroupId(todo.GroupId)
	if errUser != nil {
		response.Fail(ctx, gin.H{"error": errUser}, "group is not exits")
		return
	}
	errCount := db.DB.Model(&count).Update("item_count", gorm.Expr("item_count + 1")).Error
	if errCount != nil {
		response.Fail(ctx, gin.H{"error": errCount}, "add item failed")
		return
	}

	response.Success(ctx, gin.H{"Todo": dto.TodTodoDto(todo)}, "OK")

}

func UpdateTodo(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Fail(ctx, nil, "context error")
		return
	}
	var id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "id is not a integer value")
		return
	}
	var bind model.Todo
	err = ctx.ShouldBind(&bind)
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "bind failed")
		return
	}
	bind.TodoId = uint(id)
	bind.UserId = user.(model.User).UserId

	// this todo has a origin group, so you need to sub its count and add it to the new one
	if bind.GroupId != 0 {
		err = dao.UpdateGroupCount(bind.TodoId, bind.GroupId)
		if err != nil {
			response.Fail(ctx, gin.H{"error": err}, "update count error")
			return
		}
	}
	_, err = dao.UpdateTodo(&bind)
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "update failed")
		return
	}
	response.Success(ctx, nil, "OK")
}

func DeleteTodo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	userId := user.(model.User).UserId
	todoId, errFormat := strconv.Atoi(ctx.Param("id"))
	if errFormat != nil {
		response.Fail(ctx, gin.H{"error": errFormat}, "id is not an integer value")
		return
	}
	count, errGroup := dao.GetGroupByTodoId(uint(todoId))
	if errGroup != nil {
		response.Fail(ctx, gin.H{"error": errGroup}, "group is not exits")
		return
	}
	errDelete := dao.DeleteTodoById(uint(todoId), userId)
	if errDelete != nil {
		response.Fail(ctx, gin.H{"error": errDelete}, "delete failed")
		return
	}
	errCount := db.DB.Model(&count).Update("item_count", gorm.Expr("item_count - 1")).Error
	if errCount != nil {
		response.Fail(ctx, gin.H{"error": errCount}, "delete item failed")
		return
	}
	response.Success(ctx, nil, "OK")

}

func GetTodo(ctx *gin.Context) {
	var s dto.Condition
	errBind := ctx.ShouldBindQuery(&s)
	if s.CreatedAt.IsZero() {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		s.CreatedAt, _ = time.ParseInLocation(util.TimeFormat, time.Now().Format(util.TimeFormat), loc)
	}
	if errBind != nil {
		response.Fail(ctx, gin.H{"error": errBind.Error()}, "Bind error")
		return
	}
	todos, err := dao.GetTodoBySelectCondition(s)
	if err != nil {
		response.Fail(ctx, gin.H{"error": err.Error()}, "conditions error")
		return
	}
	response.Success(ctx, gin.H{"todos": dto.TosdTodoDto(todos)}, "OK")
}
