package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DiscussionCreateInput - input json data
type DiscussionCreateInput struct {
	UserID     uint   `json:"user_id"`
	CategoryID uint   `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// DiscussionUpdateInput - update json data
type DiscussionUpdateInput struct {
	UserID     uint   `json:"user_id"`
	CategoryID uint   `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// InsertDiscussion - insert new discussion
// POST /discussion
// Create a discussion
func InsertDiscussion(c *gin.Context) {
	var input DiscussionCreateInput
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Where("ID = ?", input.UserID).Find(&user)

	// create a discussion object
	discussion := models.Discussion{Title: input.Title, Content: input.Content, UserID: input.UserID, User: user}

	err := models.DB.Create(&discussion).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"data": discussion})
}

// FindDiscussions - retrevie all datas
// GET /discussions
// Find all discussions
func FindDiscussions(c *gin.Context) {
	var discussions []models.Discussion

	err := models.DB.Preload(clause.Associations).Find(&discussions).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": discussions})
}

// FindDiscussion - retrevie discussion by id
// GET /discussion/{id}
// Find specific discussion base on id
func FindDiscussion(c *gin.Context) {
	var discussion models.Discussion
	id := c.Param("id")

	err := models.DB.Preload(clause.Associations).Where("discussions.id = ?", id).Find(&discussion).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if discussion.UserID == 0 || discussion.Model.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": discussion})
}

// UpdateDiscussion - Update discussion
// PUT /discussion/{id}
// Update discussion base on id
func UpdateDiscussion(c *gin.Context) {
	var input DiscussionUpdateInput
	var discussion models.Discussion

	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// create discussion object
	discussion = models.Discussion{Title: input.Title, Content: input.Content, UserID: input.UserID}

	err := models.DB.Model(&discussion).Where("id = ?", id).Updates(discussion).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})
}

// DeleteDiscussion - Delete a discussion
// DELETE /discussion/{id}
// Delete a discussion base on id
func DeleteDiscussion(c *gin.Context) {
	id := c.Param("id")

	err := models.DB.Delete(&models.Discussion{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
