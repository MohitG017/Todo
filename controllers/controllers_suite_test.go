package controllers_test

import (
	"Backend_Todo/controllers"
	"Backend_Todo/initializers"
	"Backend_Todo/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestControllers runs the Ginkgo test suite
func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

// Define variables to use in the test cases
var router *gin.Engine
var db *gorm.DB

// BeforeSuite sets up the database and routes before running tests
var _ = BeforeSuite(func() {
	// Connect to an in-memory SQLite database for testing
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		Fail("Failed to connect to test database")
	}

	// Set up the database schema
	db.AutoMigrate(&models.Todo{})

	// Assign the test database to the initializers package's DB variable
	initializers.DB = db

	// Initialize a new Gin router and set up routes
	router = gin.Default()
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)
})

// AfterSuite closes the database connection
var _ = AfterSuite(func() {
	db, _ := initializers.DB.DB()
	db.Close()
})

// Test cases
var _ = Describe("TodoController", func() {
	BeforeEach(func() {
		// Clear the todos table before each test
		initializers.DB.Exec("DELETE FROM todos")
	})

	Describe("POST /todos", func() {
		It("should create a new todo", func() {
			todo := models.Todo{Title: "New Todo"}
			body, _ := json.Marshal(todo)

			req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var createdTodo models.Todo
			json.Unmarshal(w.Body.Bytes(), &createdTodo)
			Expect(createdTodo.Title).To(Equal("New Todo"))
		})
	})

	Describe("GET /todos", func() {
		It("should retrieve all todos", func() {
			initializers.DB.Create(&models.Todo{Title: "Test Todo 1"})
			initializers.DB.Create(&models.Todo{Title: "Test Todo 2"})

			req, _ := http.NewRequest("GET", "/todos", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var todos []models.Todo
			json.Unmarshal(w.Body.Bytes(), &todos)
			Expect(len(todos)).To(Equal(2))
		})
	})

	Describe("GET /todos/:id", func() {
		It("should retrieve a todo by ID", func() {
			todo := models.Todo{Title: "Single Todo"}
			initializers.DB.Create(&todo)

			// Use fmt.Sprintf to convert the ID to a string
			req, _ := http.NewRequest("GET", "/todos/"+fmt.Sprintf("%d", todo.ID), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var retrievedTodo models.Todo
			json.Unmarshal(w.Body.Bytes(), &retrievedTodo)
			Expect(retrievedTodo.Title).To(Equal("Single Todo"))
		})
	})

	Describe("PUT /todos/:id", func() {
		It("should update an existing todo", func() {
			todo := models.Todo{Title: "Old Title"}
			initializers.DB.Create(&todo)

			updatedTodo := models.Todo{Title: "Updated Title"}
			body, _ := json.Marshal(updatedTodo)

			// Use fmt.Sprintf to convert the ID to a string
			req, _ := http.NewRequest("PUT", "/todos/"+fmt.Sprintf("%d", todo.ID), bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var updated models.Todo
			json.Unmarshal(w.Body.Bytes(), &updated)
			Expect(updated.Title).To(Equal("Updated Title"))
		})
	})

	Describe("DELETE /todos/:id", func() {
		It("should delete a todo by ID", func() {
			todo := models.Todo{Title: "Delete Me"}
			initializers.DB.Create(&todo)

			// Use fmt.Sprintf to convert the ID to a string
			req, _ := http.NewRequest("DELETE", "/todos/"+fmt.Sprintf("%d", todo.ID), nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))

			var count int64
			initializers.DB.Model(&models.Todo{}).Count(&count)
			Expect(count).To(Equal(int64(0)))
		})
	})
})
