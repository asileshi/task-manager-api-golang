package routers

import (
	"github.com/gin-gonic/gin"
	"task-manager/controllers"
)

func SetupRouter() *gin.Engine {
	routers := gin.Default()
	tasks := routers.Group("/tasks")
	{
		tasks.GET("/", controllers.GetTasksHandler)
		tasks.POST("/", controllers.CreateTaskHandler)
		tasks.GET("/:id", controllers.GetTaskByIDHandler)
		tasks.PUT("/:id", controllers.UpdateTaskHandler)
		tasks.DELETE("/:id", controllers.DeleteTaskHandler)
	}
	return routers
}