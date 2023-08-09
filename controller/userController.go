package controller

import (
	"LK_blog/config/service"
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GetAllUsers(c *gin.Context, db *gorm.DB) {
	var users []model.APIUser
	db.Model(&model.User{}).Find(&users)
	c.JSON(200, users)
}

func GetUserById(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	user := model.APIUser{ID: id}
	result := db.Model(&model.User{}).First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if newUser.Birthday.IsZero() {
		c.JSON(400, gin.H{"error": "Birthday is required"})
		return
	}

	hashPassword, errHash := service.HashPassword(newUser.Password)
	if errHash != nil {
		c.JSON(400, gin.H{"error": "Hash password fail"})
	}
	newUser.ID = uuid.New()
	newUser.Password = hashPassword
	newUser.Birthday = time.Now()
	db.Create(&newUser)
	c.JSON(201, newUser)
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	user := model.User{ID: id}
	result := db.First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Parse the incoming JSON data to update the user fields
	var updateUser model.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	if updateUser.Password != "" {
		hashPassword, errHash := service.HashPassword(user.Password)
		if errHash != nil {
			c.JSON(400, gin.H{"error": "Hash password fail"})
		}
		updateUser.Password = hashPassword
	}

	// Save the updated user to the database
	result = db.Model(&user).Updates(updateUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	uuidStr := c.Param("id")
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(422, gin.H{"error": "Invalid UUID"})
		return
	}

	user := model.User{ID: id}
	result := db.First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	result = db.Delete(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
