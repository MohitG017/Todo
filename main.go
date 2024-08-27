package main

import (
	"Backend_Todo/initializers"
	"Backend_Todo/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize the database connection
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080") // Default port is 8080
}
