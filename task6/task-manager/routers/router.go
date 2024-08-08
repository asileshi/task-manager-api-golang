package routers

import (
	"github.com/gin-gonic/gin"
	"task-manager/controllers"
	"task-manager/middleware"
)

func SetupRouter() *gin.Engine {
	routers := gin.Default()
	users := routers.Group("/users")
	{
		users.POST("/register", controllers.RegistrationHandler)
		users.POST("/login", controllers.LoginHandler)
		users.PUT("/promote/:id", middleware.AuthMiddleware, middleware.AdminMidleware, controllers.PromotoionHandler)
	}

	tasks := routers.Group("/tasks")
	{
		tasks.GET("/", controllers.GetTasksHandler)
		tasks.GET("/:id", controllers.GetTaskByIDHandler)

		tasks.POST("/", middleware.AuthMiddleware, middleware.AdminMidleware, controllers.CreateTaskHandler)
		tasks.PUT("/:id", middleware.AuthMiddleware, middleware.AdminMidleware, controllers.UpdateTaskHandler)
		tasks.DELETE("/:id", middleware.AuthMiddleware, middleware.AdminMidleware,controllers.DeleteTaskHandler)
	}

	return routers
}