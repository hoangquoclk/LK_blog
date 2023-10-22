package router

import (
	"LK_blog/controller"
	"LK_blog/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", middleware.Authentication, func(context *gin.Context) {
		controller.GetAllUsers(context, db)
	})
	r.GET("/users/:id", middleware.Authentication, func(context *gin.Context) {
		controller.GetUserById(context, db)
	})
	r.POST("/users", middleware.Authentication, func(context *gin.Context) {
		controller.CreateUser(context, db)
	})
	r.POST("/users/login", func(context *gin.Context) {
		controller.Login(context, db)
	})
	r.PUT("/users/:id", middleware.Authentication, func(context *gin.Context) {
		controller.UpdateUser(context, db)
	})
	r.DELETE("/users/:id", middleware.Authentication, func(context *gin.Context) {
		controller.DeleteUser(context, db)
	})
}
