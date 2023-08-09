package controller

import (
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllCategories(c *gin.Context, db *gorm.DB) {
	var categories []model.Category
	db.Find(&categories)
	c.JSON(200, categories)
}

func GetCategoryById(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	category := model.Category{ID: id}
	result := db.First(&category)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "category not found"})
		return
	}

	c.JSON(200, category)
}

func CreateCategory(c *gin.Context, db *gorm.DB) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	newCategory.ID = uuid.New()
	db.Create(&newCategory)
	c.JSON(201, newCategory)
}

func UpdateCategory(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	category := model.Category{ID: id}
	result := db.First(&category)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	// Parse the incoming JSON data to update the category fields
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Save the updated category to the database
	result = db.Save(&category)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(200, category)
}

func DeleteCategory(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	category := model.Category{ID: id}
	result := db.First(&category)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	// Delete the category from the database
	result = db.Delete(&category)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}
