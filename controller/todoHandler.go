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
	user, _ := ctx.Get("user")
	id, errFormat := strconv.Atoi(ctx.Param("id"))
	if errFormat != nil {
		response.Fail(ctx, gin.H{"error": errFormat}, "id is not a integer value")
		return
	}
	var bind model.Todo
	errBind := ctx.ShouldBind(&bind)
	if errBind != nil {
		response.Fail(ctx, gin.H{"error": errBind.Error()}, "bind failed")
		return
	}
	bind.TodoId = uint(id)
	bind.UserId = user.(model.User).UserId
	if bind.GroupId != 0 {
		origin, errBefore := dao.GetGroupByTodoId(bind.TodoId)
		after, errAfter := dao.GetGroupByGroupId(bind.GroupId)
		if errBefore != nil {
			response.Fail(ctx, gin.H{"error": errBefore}, "original group is not exits")
			return
		} else if errAfter != nil {
			response.Fail(ctx, gin.H{"error": errAfter}, "after group is not exits")
			return
		}
		err_original := db.DB.Model(&origin).Update("item_count", gorm.Expr("item_count - 1")).Error
		err_after := db.DB.Model(&after).Update("item_count", gorm.Expr("item_count + 1")).Error
		if err_original != nil || err_after != nil {
			response.Fail(ctx, gin.H{"error_or": err_original, "error_ad": err_after}, "update items error")
			return
		}
		_, errUpdate := dao.UpdateTodo(&bind)
		if errUpdate != nil {
			response.Fail(ctx, gin.H{"error": errUpdate}, "update failed")
			return
		}
		response.Success(ctx, nil, "OK")
		return
	}
	_, errUpdate := dao.UpdateTodo(&bind)
	if errUpdate != nil {
		response.Fail(ctx, gin.H{"error": errUpdate}, "update failed")
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
