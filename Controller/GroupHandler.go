package Controller

import (
	"github.com/WenkanHuang/gin_gorm/Dao"
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/Response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddGroup(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	groupName := ctx.PostForm("groupName")
	userId := user.(Model.User).UserId
	group, err := Dao.AddGroup(&Model.Group{
		GroupName: groupName,
		UserId:    userId,
	})
	if err != nil {
		Response.Fail(ctx, nil, "add failed")
	} else {
		Response.Success(ctx, gin.H{"group": group}, "OK")
	}
}

func ShowGroupList(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	groupList, err := Dao.GetGroupsByUserId(user.(Model.User).UserId)
	if err != nil {
		Response.Fail(ctx, nil, "add failed")
	} else {
		Response.Success(ctx, gin.H{"todo_group_list": groupList}, "OK")
	}
}
func UpdateGroup(ctx *gin.Context) {
	groupName := ctx.PostForm("groupName")
	id, errInString := strconv.Atoi(ctx.Param("id"))
	if errInString != nil {
		Response.Fail(ctx, gin.H{"error": errInString}, "parameter error")
	}
	group, err := Dao.UpdateGroup(&Model.Group{
		GroupId:   uint(id),
		GroupName: groupName,
	})
	if err != nil {
		Response.Fail(ctx, gin.H{"error": err}, "update failed")
	} else {
		Response.Success(ctx, gin.H{"group": group}, "OK")
	}
}

func DeleteGroupById(ctx *gin.Context) {
	id, errInString := strconv.Atoi(ctx.Param("id"))
	if errInString != nil {
		Response.Fail(ctx, gin.H{"error": errInString}, "parameter error")
	}
	err := Dao.DeleteGroupById(uint(id))
	if err != nil {
		Response.Fail(ctx, gin.H{"error": err}, "update failed")
	} else {
		Response.Success(ctx, nil, "OK")
	}
}
