package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateCommentInput - input comment json
type CreateCommentInput struct {
	Content      string `json:"content"`
	ParentID     uint   `json:"parent_id"`
	UserID       uint   `json:"user_id"`
	DiscussionID uint   `json:"discussion_id"`
}

// UpdateCommentInput - update comment json
type UpdateCommentInput struct {
	Content      string `json:"content"`
	ParentID     uint   `json:"parent_id"`
	UserID       uint   `json:"user_id"`
	DiscussionID uint   `json:"discussion_id"`
}

// InsertComment - insert new comment
// POST /comment
// Insert new comment related to a Comment
func InsertComment(c *gin.Context) {
	var comment models.Comment
	var input CreateCommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comment = models.Comment{ParentID: input.ParentID, Content: input.Content, UserID: input.UserID, DiscussionID: input.DiscussionID}

	// insert into database
	err := models.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// FindComments - retrieve all comments
// GET /comments
// Find all comments
func FindComments(c *gin.Context) {
	var comments []models.Comment
	models.DB.Find(&comments)
	// err := models.DB.Joins("User").Find(&comments).Error
	err := models.DB.Preload(clause.Associations).Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// FindComment - retrieve all comments
// GET /comment
// Find all comment
func FindComment(c *gin.Context) {
	var comments models.Comment
	id := c.Param("id")

	err := models.DB.Find(&comments, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if comments.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateComment - Update comment
// PUT /comment/{id}
// Update comment base on id
func UpdateComment(c *gin.Context) {
	var input UpdateCommentInput
	var comment models.Comment

	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// create Comment object
	comment = models.Comment{ParentID: input.ParentID, Content: input.Content, UserID: input.UserID, DiscussionID: input.DiscussionID}

	err := models.DB.Model(&comment).Where("id = ?", id).Updates(comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})
}

// DeleteComment - Delete a Comment
// DELETE /Comment/{id}
// Delete a Comment base on id
func DeleteComment(c *gin.Context) {
	id := c.Param("id")

	err := models.DB.Delete(&models.Comment{}, id).Error
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
