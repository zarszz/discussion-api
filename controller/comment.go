package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
)

// GET /comments
// Find all comments
func FindComments(c *gin.Context) {
	var comments []models.Comment
	models.DB.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}
