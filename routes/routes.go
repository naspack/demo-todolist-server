package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naspack/demo-todolist-server/controllers"
)

func SetupRoutes(r *gin.Engine) {
	// API路由组
	api := r.Group("/api")

	// Todo相关路由
	todos := api.Group("/todos")
	{
		todos.POST("", controllers.CreateTodo)
		todos.GET("", controllers.GetTodos)
		todos.GET("/:id", controllers.GetTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}
}