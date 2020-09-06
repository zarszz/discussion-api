package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
)

type CreateUserInput struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhotoProfile string `json:"photo_profile"`
}

// GET /users
// Find all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /user
// Create a new user
func InsertUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create user
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, PhotoProfile: input.PhotoProfile}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
