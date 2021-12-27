package Middleware

import (
	"net/http"
	"strings"

	"github.com/WenkanHuang/gin_gorm/Common"
	"github.com/WenkanHuang/gin_gorm/Db"
	"github.com/WenkanHuang/gin_gorm/Model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization") // 获取authorization header
		prefix := "Bearer"

		//validate token
		if tokenString == "" || !strings.HasPrefix(tokenString, prefix) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "Auth Method is wrong"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[len(prefix)+1:]
		token, claims, err := Common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		userId := claims.ID // 验证通过之后获取claim中的userId
		var user Model.User
		Db.DB.First(&user, userId)
		if user.UserId == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "msg": "权限不够"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
