package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zarszz/discussion-api/models"
	"gorm.io/gorm"
)

type CreateCategoryInput struct {
	Category string `json:"category"`
}

type CategoryUpdateInput struct {
	Category string `json:"category"`
}

// FindCategories - retrieve all category
// GET /categories
// Find all users
func FindCategories(c *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// FindCategory - retrevie discussion by id
// GET /category/{id}
// Find specific category base on id
func FindCategory(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	// err := models.DB.Joins("Comment").First(&discussion, "discussions.id = ?", id).Error
	err := models.DB.Where("categories.id = ?", id).Find(&category).Error
	// err := models.DB.Where("id = ?", id).Preload("Comment").Find(&discussion).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// InsertCategory - create new category
// POST /Category
// Create a new Category
func InsertCategory(c *gin.Context) {
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create category

	category := models.Category{Category: input.Category}
	models.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory - Update Category
// PUT /category/{id}
// Update Category base on id
func UpdateCategory(c *gin.Context) {
	var input CategoryUpdateInput
	var category models.Category

	id := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// create Category object
	category = models.Category{Category: input.Category}

	err := models.DB.Model(&category).Where("id = ?", id).Updates(category).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "update successfully"})
}

// DeleteCategory - Delete a Category
// DELETE /category/{id}
// Delete a Category base on id
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	err := models.DB.Delete(&models.Category{}, id).Error
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
