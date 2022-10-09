package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/WenkanHuang/gin_gorm/dao"
	"github.com/WenkanHuang/gin_gorm/dto"
	"github.com/WenkanHuang/gin_gorm/model"
	"github.com/WenkanHuang/gin_gorm/response"
)

func AddGroup(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	groupName := ctx.PostForm("groupName")
	userId := user.(model.User).UserId
	group, err := dao.AddGroup(&model.Group{
		GroupName: groupName,
		UserId:    userId,
	})
	if err != nil {
		log.Errorf("Add gruop error: %+v", errors.WithStack(err))
		response.Fail(ctx, nil, "add failed")
		return
	}
	response.Success(ctx, gin.H{"group": dto.ToGroupDto(*group)}, "OK")
}

func ShowGroupList(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		log.Errorf("user does not foud")
		response.Fail(ctx, nil, "context does not have user field error")
		return
	}
	groupList, err := dao.GetGroupsByUserId(user.(model.User).UserId)
	if err != nil {
		response.Fail(ctx, nil, "add failed")
		return
	}
	response.Success(ctx, gin.H{"todo_group_list": groupList}, "OK")

}

func UpdateGroup(ctx *gin.Context) {
	groupName := ctx.PostForm("groupName")
	id, errInString := strconv.Atoi(ctx.Param("id"))
	if errInString != nil {
		response.Fail(ctx, gin.H{"error": errInString}, "parameter error")
	}
	group, err := dao.UpdateGroup(&model.Group{
		GroupId:   uint(id),
		GroupName: groupName,
	})
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "update failed")
	} else {
		response.Success(ctx, gin.H{"group": dto.ToGroupDto(*group)}, "OK")
	}
}

func DeleteGroupById(ctx *gin.Context) {
	id, errInString := strconv.Atoi(ctx.Param("id"))
	if errInString != nil {
		response.Fail(ctx, gin.H{"error": errInString}, "parameter error")
	}
	err := dao.DeleteGroupById(uint(id))
	if err != nil {
		response.Fail(ctx, gin.H{"error": err}, "update failed")
	} else {
		response.Success(ctx, nil, "OK")
	}
}
