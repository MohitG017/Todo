package controllers

import (
	"Backend_Todo/initializers"
	"Backend_Todo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTodo creates a new todo item
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetTodos retrieves all todo items
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := initializers.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetTodo retrieves a todo item by ID
func GetTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo updates an existing todo item
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo item by ID
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	initializers.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
