package controller

import (
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context, db *gorm.DB) {
	var posts []model.Post
	db.Find(&posts)
	c.JSON(200, posts)
}

func GetPostById(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	post := model.Post{ID: id}
	result := db.First(&post)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "post not found"})
		return
	}

	c.JSON(200, post)
}

func CreatePost(c *gin.Context, db *gorm.DB) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}
	newPost.ID = uuid.New()
	db.Create(&newPost)
	c.JSON(201, newPost)
}

func UpdatePost(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	post := model.Post{ID: id}
	result := db.First(&post)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// Parse the incoming JSON data to update the post fields
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	// Save the updated post to the database
	result = db.Save(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(200, post)
}

func DeletePost(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	post := model.Post{ID: id}
	result := db.First(&post)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	// Delete the post from the database
	result = db.Delete(&post)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
