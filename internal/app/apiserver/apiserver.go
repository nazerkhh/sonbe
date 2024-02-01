// internal/app/apiserver/apiserver.go
package apiserver

import (
	"github.com/Nazerkh09/sonbe/internal/app/apiserver/handlers"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// User authentication routes
	r.POST("/api/register", handlers.RegisterUser)
	r.POST("/api/login", handlers.LoginUser)
	r.POST("/api/logout", handlers.LogoutUser)

	// Group management routes
	r.POST("/api/groups", handlers.CreateGroup)
	r.GET("/api/groups/:id", handlers.GetGroup)
	r.PUT("/api/groups/:id", handlers.UpdateGroup)
	r.DELETE("/api/groups/:id", handlers.DeleteGroup)

	// Task management routes
	r.POST("/api/tasks", handlers.CreateTask)
	r.GET("/api/tasks/:id", handlers.GetTask)
	r.PUT("/api/tasks/:id", handlers.UpdateTask)
	r.DELETE("/api/tasks/:id", handlers.DeleteTask)

	// DOMjudge integration routes
	r.POST("/api/submit", handlers.SubmitToDOMjudge)

	r.Run(":8080")
}
