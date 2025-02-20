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
		todos.POST("/get", controllers.GetTodo)
		todos.POST("/update", controllers.UpdateTodo)
		todos.POST("/delete", controllers.DeleteTodo)
		todos.POST("/complete", controllers.CompleteTodo)
		todos.POST("/uncomplete", controllers.UncompleteTodo)
	}
}
