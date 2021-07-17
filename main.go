package main

import (
	"github.com/WenkanHuang/gin_gorm/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", Controller.Ping)
	r.Run()

}
