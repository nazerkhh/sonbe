// internal/app/apiserver/routes/routes.go
package routes

import (
	"github.com/Nazerkhh/sonbe/internal/app/apiserver/handlers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the API
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// User routes
	api.POST("/register", handlers.RegisterUser)
	api.POST("/login", handlers.LoginUser)
	api.POST("/logout", handlers.LogoutUser)

	// Group routes
	api.POST("/groups", handlers.CreateGroup)
	api.GET("/groups/:id", handlers.GetGroup)
	api.PUT("/groups/:id", handlers.UpdateGroup)
	api.DELETE("/groups/:id", handlers.DeleteGroup)

	// Task routes
	api.POST("/tasks", handlers.CreateTask)
	api.GET("/tasks/:id", handlers.GetTask)
	api.PUT("/tasks/:id", handlers.UpdateTask)
	api.DELETE("/tasks/:id", handlers.DeleteTask)

	// DOMjudge submission route
	api.POST("/submit", handlers.SubmitToDOMjudge)
}
