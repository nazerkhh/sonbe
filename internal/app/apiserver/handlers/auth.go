// internal/app/apiserver/handlers/auth.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User struct represents a basic user model
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User

func init() {
	// Pre-populate some dummy users for demonstration purposes
	users = append(users, User{ID: 1, Username: "user1", Password: "password1"})
	users = append(users, User{ID: 2, Username: "user2", Password: "password2"})
}

// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username is already taken
	for _, existingUser := range users {
		if existingUser.Username == newUser.Username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}
	}

	// Assign a new ID and add the user to the list
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginUser handles user login
func LoginUser(c *gin.Context) {
	var credentials User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user with the provided username and password
	for _, existingUser := range users {
		if existingUser.Username == credentials.Username && existingUser.Password == credentials.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// LogoutUser handles user logout
func LogoutUser(c *gin.Context) {
	// Perform any necessary logout logic (if needed)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
