/*
 * @Auther: Edward
 * @Date: 2021-07-19 22:12:14
 * @LastEditTime: 2021-08-29 17:17:00
 * @FilePath: \gin_gorm\Response\apiResponsse.go
 */
package Response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
