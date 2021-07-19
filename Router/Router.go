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
		user.POST("/register", Controller.Register)
		user.POST("/login", Controller.Login)
		user.GET("/info", Middleware.AuthMiddleware(), Controller.Info)
	}

	group := r.Group("/v1/group")
	{
		group.POST("/add", Middleware.AuthMiddleware(), Controller.AddGroup)
		group.PUT("/:id", Middleware.AuthMiddleware(), Controller.UpdateGroup)
		group.DELETE("/:id", Middleware.AuthMiddleware(), Controller.DeleteGroupById)
		group.GET("/list", Middleware.AuthMiddleware(), Controller.ShowGroupList)
	}

	//
	//postRoutes := r.Group("/post")
	//postRoutes.Use(Middleware.AuthMiddleware())
	//postController := Controller.NewPostController()
	//postRoutes.POST("/add", postController.AddPost)
	//postRoutes.GET("/getAll", postController.GetPost)
	//postRoutes.PUT("/update", postController.UpdatePost)
	//postRoutes.DELETE("/delete", postController.DeletePostByID)

	return r
}
