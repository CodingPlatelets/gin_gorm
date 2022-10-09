package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/WenkanHuang/gin_gorm/common"
	"github.com/WenkanHuang/gin_gorm/db"
	"github.com/WenkanHuang/gin_gorm/dto"
	"github.com/WenkanHuang/gin_gorm/model"
	"github.com/WenkanHuang/gin_gorm/response"
	"github.com/WenkanHuang/gin_gorm/util"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	if len(password) < 6 {
		response.SimpleResponse(ctx, http.StatusUnprocessableEntity, nil, "Password less than 6 digits!")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		response.SimpleResponse(ctx, http.StatusInternalServerError, nil, "Hashed password error!")
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hashedPassword),
	}
	db.DB.Create(&user)
	response.Success(ctx, nil, "Register success!")
}

func Login(ctx *gin.Context) {
	name := ctx.PostForm("name") // 获取参数
	password := ctx.PostForm("password")
	if password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password not null!"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password len not less 6 bits!"})
		return
	}
	var user model.User //判断是否存在
	db.DB.Where("name = ?", name).First(&user)
	if user.UserId == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User not exist!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // 判断密码是否正确
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password err!"})
		return
	}
	token, err := common.ReleaseToken(user) // 发放token
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
		log.Printf("token generate error:%v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
