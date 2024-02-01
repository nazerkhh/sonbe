// internal/app/apiserver/handlers/group.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Group struct represents a basic group model
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var groups []Group

// CreateGroup handles group creation
func CreateGroup(c *gin.Context) {
	var newGroup Group
	if err := c.ShouldBindJSON(&newGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new ID and add the group to the list
	newGroup.ID = len(groups) + 1
	groups = append(groups, newGroup)

	c.JSON(http.StatusCreated, newGroup)
}

// GetGroup handles group retrieval
func GetGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	// Find the group with the provided ID
	for _, existingGroup := range groups {
		if existingGroup.ID == groupID {
			c.JSON(http.StatusOK, existingGroup)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
}

// UpdateGroup handles group update
func UpdateGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	var updatedGroup Group
	if err := c.ShouldBindJSON(&updatedGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the group with the provided ID and update its properties
	for i, existingGroup := range groups {
		if existingGroup.ID == groupID {
			groups[i] = updatedGroup
			c.JSON(http.StatusOK, updatedGroup)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
}

// DeleteGroup handles group deletion
func DeleteGroup(c *gin.Context) {
	groupID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	// Remove the group with the provided ID from the list
	for i, existingGroup := range groups {
		if existingGroup.ID == groupID {
			groups = append(groups[:i], groups[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
}
