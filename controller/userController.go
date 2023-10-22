package controller

import (
	"LK_blog/config/service"
	"LK_blog/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"regexp"
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

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
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

	if newUser.Email == "" {
		c.JSON(400, gin.H{"error": "Email is required"})
		return
	}

	var validMail = isValidEmail(newUser.Email)
	if validMail {
		service.SendMail(newUser)
	} else {
		c.JSON(400, gin.H{"error": "Invalid email format, please correct your email!"})
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

	result = db.Delete(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func Login(c *gin.Context, db *gorm.DB) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	var existingUser model.User

	db.Model(&model.User{Email: user.Email}).First(&existingUser)

	if existingUser.ID == uuid.Nil {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := service.CheckPasswordHash(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	tokenString, err, expirationTime := service.GenerateToken(existingUser.Role, existingUser.Email)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"token": tokenString, "expiresAt": expirationTime.Unix()})
}
