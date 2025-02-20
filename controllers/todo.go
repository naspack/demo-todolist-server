package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naspack/demo-todolist-server/config"
	"github.com/naspack/demo-todolist-server/models"
)

type TodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type TodoIDRequest struct {
	ID uint `json:"id" binding:"required"`
}

type TodoUpdateRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

func CreateTodo(c *gin.Context) {
	var req TodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	todo := models.Todo{
		Title:       req.Title,
		Description: req.Description,
	}

	if err := config.DB.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "创建待办事项失败"})
		return
	}

	c.JSON(201, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := config.DB.Find(&todos).Error; err != nil {
		c.JSON(500, gin.H{"error": "获取待办事项列表失败"})
		return
	}

	c.JSON(200, todos)
}

func GetTodo(c *gin.Context) {
	var req TodoIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, req.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "待办事项不存在"})
		return
	}

	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context) {
	var req TodoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, req.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "待办事项不存在"})
		return
	}

	todo.Title = req.Title
	todo.Description = req.Description
	if req.Completed != nil {
		todo.Completed = *req.Completed
	}

	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新待办事项失败"})
		return
	}

	c.JSON(200, todo)
}

func DeleteTodo(c *gin.Context) {
	var req TodoIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	result := config.DB.Delete(&models.Todo{}, req.ID)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "删除待办事项失败"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "待办事项不存在"})
		return
	}

	c.Status(204)
}

func CompleteTodo(c *gin.Context) {
	var req TodoIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, req.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "待办事项不存在"})
		return
	}

	todo.Completed = true
	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新待办事项失败"})
		return
	}

	c.JSON(200, todo)
}

func UncompleteTodo(c *gin.Context) {
	var req TodoIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "无效的请求参数"})
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, req.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "待办事项不存在"})
		return
	}

	todo.Completed = false
	if err := config.DB.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新待办事项失败"})
		return
	}

	c.JSON(200, todo)
}
