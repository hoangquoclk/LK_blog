package controller

import (
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllComments(c *gin.Context, db *gorm.DB) {
	var comments []model.Comment
	db.Find(&comments)
	c.JSON(200, comments)
}

func GetCommentById(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	comment := model.Comment{ID: id}
	result := db.First(&comment)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "comment not found"})
		return
	}

	c.JSON(200, comment)
}

func CreateComment(c *gin.Context, db *gorm.DB) {
	var newComment model.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	if newComment.Content == "" {
		c.JSON(400, gin.H{"error": "Content is required"})
		return
	}
	newComment.ID = uuid.New()
	db.Create(&newComment)
	c.JSON(201, newComment)
}

func UpdateComment(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	comment := model.Comment{ID: id}
	result := db.First(&comment)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	// Parse the incoming JSON data to update the comment fields
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if comment.Content == "" {
		c.JSON(400, gin.H{"error": "Content is required"})
		return
	}

	// Save the updated comment to the database
	result = db.Save(&comment)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(200, comment)
}

func DeleteComment(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	comment := model.Comment{ID: id}
	result := db.First(&comment)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Comment not found"})
		return
	}

	// Delete the comment from the database
	result = db.Delete(&comment)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}
