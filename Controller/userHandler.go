package Controller

import (
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Response"
	"github.com/WenkanHuang/gin_gorm/Util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	if len(phone) != 11 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "The phone num must be 11 digits!")
		return
	}
	if len(password) < 6 {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "Password not less 6 digits!")
		return
	}
	if len(name) == 0 {
		name = Util.RandomString(10)
	}
	if dao.IsPhoneExist(phone) {
		Response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "User exisit!")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // 创建用户的时候要加密用户的密码
	if err != nil {
		Response.Response(ctx, http.StatusInternalServerError, 500, nil, "Hashed password error!")
		return
	}
	user := model.User{
		Name:     name,
		Password: string(hashedPassword),
		Phone:    phone,
	}
	Db.DB.Create(&user)
	Response.Succces(ctx, nil, "Register success!")
}

func Login(ctx *gin.Context) {
	phone := ctx.PostForm("phone") // 获取参数
	password := ctx.PostForm("password")
	if phone == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num not null!"})
		return
	}
	if password == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password not null!"})
		return
	}
	if len(phone) != 11 { // 数据验证
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Phone num must 11 digits!"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "Password len not less 6 bigits!"})
		return
	}
	var user model.User // 判断手机号是否存在
	Db.DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "User not exist!"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // 判断密码是否正确
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "password err!"})
	}
	token, err := common.ReleaseToken(user) // 发放token
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "System err!"})
		log.Printf("token generate error:%v", err)
		return
	}
	Response.Succces(ctx, gin.H{"token": token}, "Login success!")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}
