package routes

import (
	"Backend_Todo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Todo routes
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
}
