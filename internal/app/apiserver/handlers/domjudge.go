// internal/app/apiserver/handlers/domjudge.go
package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const domjudgeAPIURL = "http://domjudge.example.com/api/"

// SubmitToDOMjudge handles code submission to DOMjudge
func SubmitToDOMjudge(c *gin.Context) {
	var submissionData struct {
		GroupID  int    `json:"group_id"`
		TaskID   int    `json:"task_id"`
		Language string `json:"language"`
		Code     string `json:"code"`
	}

	if err := c.ShouldBindJSON(&submissionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare submission payload
	submissionPayload := fmt.Sprintf(`{
		"group_id": %d,
		"task_id": %d,
		"language": "%s",
		"source_code": "%s"
	}`, submissionData.GroupID, submissionData.TaskID, submissionData.Language, submissionData.Code)

	// Perform HTTP POST request to DOMjudge API
	response, err := http.Post(domjudgeAPIURL+"submissions", "application/json", bytes.NewBufferString(submissionPayload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit code to DOMjudge"})
		return
	}
	defer response.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read DOMjudge API response"})
		return
	}

	// Check if the submission was successful based on the API response
	if response.StatusCode == http.StatusCreated {
		c.JSON(http.StatusCreated, gin.H{"message": "Code submitted to DOMjudge successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": string(responseBody)})
	}
}
