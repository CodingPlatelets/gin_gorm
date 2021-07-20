package Router

import (
	"github.com/WenkanHuang/gin_gorm/Controller"
	"github.com/WenkanHuang/gin_gorm/Middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Middleware.CORSMiddleware(), Middleware.RecoverMiddleware())

	user := r.Group("/v1/user")
	{
		user.POST("/register", Middleware.CORSMiddleware(), Controller.Register)
		user.POST("/login", Middleware.CORSMiddleware(), Controller.Login)
		user.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
	}

	group := r.Group("/v1/group")
	{
		group.POST("/add", Middleware.AuthMiddleware(), Controller.AddGroup)
		group.PUT("/:id", Middleware.AuthMiddleware(), Controller.UpdateGroup)
		group.DELETE("/:id", Middleware.AuthMiddleware(), Controller.DeleteGroupById)
		group.GET("/list", Middleware.AuthMiddleware(), Controller.ShowGroupList)
	}

	todo := r.Group("/v1/todo")
	{
		todo.POST("/add", Middleware.AuthMiddleware(), Controller.AddTodo)
		todo.PUT("/:id", Middleware.AuthMiddleware(), Controller.UpdateTodo)
		todo.DELETE("/:id", Middleware.AuthMiddleware(), Controller.DeleteTodo)
		todo.GET("/list", Middleware.AuthMiddleware(), Controller.GetTodo)
	}

	return r
}
