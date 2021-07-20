package Controller

import (
	"github.com/WenkanHuang/gin_gorm/Dao"
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Dto"
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/Response"
	"github.com/WenkanHuang/gin_gorm/Util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
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
		errCount := Db.DB.Model(&count).Update("item_count", gorm.Expr("item_count + 1")).Error
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
			if bind.GroupId != 0 {
				origin, errBefore := Dao.GetGroupByTodoId(bind.TodoId)
				after, errAfter := Dao.GetGroupByGroupId(bind.GroupId)
				if errBefore != nil {
					Response.Fail(ctx, gin.H{"error": errBefore}, "original group is not exits")
					return
				} else if errAfter != nil {
					Response.Fail(ctx, gin.H{"error": errAfter}, "after group is not exits")
					return
				} else {
					err_original := Db.DB.Model(&origin).Update("item_count", gorm.Expr("item_count - 1")).Error
					err_after := Db.DB.Model(&after).Update("item_count", gorm.Expr("item_count + 1")).Error
					if err_original != nil || err_after != nil {
						Response.Fail(ctx, gin.H{"error_or": err_original, "error_ad": err_after}, "update items error")
						return
					} else {
						_, errUpdate := Dao.UpdateTodo(&bind)
						if errUpdate != nil {
							Response.Fail(ctx, gin.H{"error": errUpdate}, "update failed")
							return
						} else {
							Response.Success(ctx, nil, "OK")
							return
						}
					}
				}
			}
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
		return
	} else {
		count, errGroup := Dao.GetGroupByTodoId(uint(todoId))
		if errGroup != nil {
			Response.Fail(ctx, gin.H{"error": errGroup}, "group is not exits")
			return
		} else {
			errDelete := Dao.DeleteTodoById(uint(todoId), userId)
			if errDelete != nil {
				Response.Fail(ctx, gin.H{"error": errDelete}, "delete failed")
				return
			} else {
				errCount := Db.DB.Model(&count).Update("item_count", gorm.Expr("item_count - 1")).Error
				if errCount != nil {
					Response.Fail(ctx, gin.H{"error": errCount}, "delete item failed")
					return
				}
			}
			Response.Success(ctx, nil, "OK")
		}

	}
}

func GetTodo(ctx *gin.Context) {
	var s Dto.Condition
	errBind := ctx.ShouldBindQuery(&s)
	if s.CreatedAt.IsZero() {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		s.CreatedAt, _ = time.ParseInLocation(Util.TimeFormat, time.Now().Format(Util.TimeFormat), loc)
	}
	if errBind != nil {
		Response.Fail(ctx, gin.H{"error": errBind.Error()}, "Bind error")
	} else {
		todos, err := Dao.GetTodoBySelectCondition(s)
		if err != nil {
			Response.Fail(ctx, gin.H{"error": err.Error()}, "conditions error")
		} else {
			Response.Success(ctx, gin.H{"todos": Dto.TosdTodoDto(todos)}, "OK")
		}
	}

}
