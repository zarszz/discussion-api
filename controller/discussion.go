package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
)

// GET /discussions
// Find all discussions
func FindDiscussions(c *gin.Context) {
	var discussions []models.Discussion
	models.DB.Find(&discussions)

	c.JSON(http.StatusOK, gin.H{"data": discussions})
}
