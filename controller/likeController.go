package controller

import (
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllLikes(c *gin.Context, db *gorm.DB) {
	var likes []model.Like
	db.Find(&likes)
	c.JSON(200, likes)
}

func GetLikeById(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	like := model.Like{ID: id}
	result := db.First(&like)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "like not found"})
		return
	}

	c.JSON(200, like)
}

func CreateLike(c *gin.Context, db *gorm.DB) {
	var newLike model.Like
	if err := c.ShouldBindJSON(&newLike); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	newLike.ID = uuid.New()
	db.Create(&newLike)
	c.JSON(201, newLike)
}

func UpdateLike(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	like := model.Like{ID: id}
	result := db.First(&like)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Like not found"})
		return
	}

	// Parse the incoming JSON data to update the like fields
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Save the updated like to the database
	result = db.Save(&like)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update like"})
		return
	}

	c.JSON(200, like)
}

func DeleteLike(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID"})
		return
	}

	like := model.Like{ID: id}
	result := db.First(&like)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Like not found"})
		return
	}

	// Delete the like from the database
	result = db.Delete(&like)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete like"})
		return
	}

	c.JSON(200, gin.H{"message": "Like deleted successfully"})
}
