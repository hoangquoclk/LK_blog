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

func CreateLike(c *gin.Context, db *gorm.DB) {
	var newLike model.Like
	if err := c.ShouldBindJSON(&newLike); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	db.Create(&newLike)
	c.JSON(201, newLike)
}

func DeleteLike(c *gin.Context, db *gorm.DB) {
	userIdStr := c.Param("userId")
	postIdStr := c.Param("postId")
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UserId"})
		return
	}
	postId, postErr := uuid.Parse(postIdStr)
	if postErr != nil {
		c.JSON(422, gin.H{"error": "Invalid PostID"})
		return
	}

	like := model.Like{UserId: userId, PostId: postId}
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
