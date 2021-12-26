/*
 * @Auther: Platelets
 * @Date: 2021-11-10 16:29:55
 * @LastEditTime: 2021-12-10 16:18:36
 * @FilePath: \gin_gorm\Controller\UserHandler.go
 */
package Controller

import (
	"log"
	"net/http"

	"github.com/WenkanHuang/gin_gorm/Common"
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Dto"
	"github.com/WenkanHuang/gin_gorm/Model"
	"github.com/WenkanHuang/gin_gorm/Response"
	"github.com/WenkanHuang/gin_gorm/Util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	if len(password) < 6 {
		Response.SimpleResponse(ctx, http.StatusUnprocessableEntity, nil, "Password not less 6 digits!")
		return
	}
	if len(name) == 0 {
		name = Util.RandomString(10)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		Response.SimpleResponse(ctx, http.StatusInternalServerError, nil, "Hashed password error!")
		return
	}
	user := Model.User{
		Name:     name,
		Password: string(hashedPassword),
	}
	Db.DB.Create(&user)
	Response.Success(ctx, nil, "Register success!")
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
	var user Model.User //判断是否存在
	Db.DB.Where("name = ?", name).First(&user)
	if user.UserId == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User not exist!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // 判断密码是否正确
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password err!"})
	}
	token, err := Common.ReleaseToken(user) // 发放token
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
		log.Printf("token generate error:%v", err)
		return
	}
	Response.Success(ctx, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": Dto.ToUserDto(user.(Model.User))}})
}
