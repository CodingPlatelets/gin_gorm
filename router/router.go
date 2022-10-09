package router

import (
	"github.com/gin-gonic/gin"

	"github.com/WenkanHuang/gin_gorm/controller"
	"github.com/WenkanHuang/gin_gorm/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware(), middleware.RecoverMiddleware())
	r.GET("/ping", controller.Ping)
	user := r.Group("/v1/user")
	{
		user.POST("/register", middleware.CORSMiddleware(), controller.Register)
		user.POST("/login", middleware.CORSMiddleware(), controller.Login)
		user.GET("/info", middleware.AuthMiddleware(), controller.Info)
	}

	group := r.Group("/v1/group")
	{
		group.POST("/add", middleware.AuthMiddleware(), controller.AddGroup)
		group.PUT("/:id", middleware.AuthMiddleware(), controller.UpdateGroup)
		group.DELETE("/:id", middleware.AuthMiddleware(), controller.DeleteGroupById)
		group.GET("/list", middleware.AuthMiddleware(), controller.ShowGroupList)
	}

	todo := r.Group("/v1/todo")
	{
		todo.POST("/add", middleware.AuthMiddleware(), controller.AddTodo)
		todo.PUT("/:id", middleware.AuthMiddleware(), controller.UpdateTodo)
		todo.DELETE("/:id", middleware.AuthMiddleware(), controller.DeleteTodo)
		todo.GET("/list", middleware.AuthMiddleware(), controller.GetTodo)
	}

	return r
}
