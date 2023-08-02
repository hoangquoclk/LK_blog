package router

import (
	"LK_blog/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", func(context *gin.Context) {
		controller.GetAllUsers(context, db)
	})
	r.GET("/users/:id", func(context *gin.Context) {
		controller.GetUserById(context, db)
	})
	r.POST("/users", func(context *gin.Context) {
		controller.CreateUser(context, db)
	})
	r.PUT("/users/:id", func(context *gin.Context) {
		controller.UpdateUser(context, db)
	})
	r.DELETE("/users/:id", func(context *gin.Context) {
		controller.DeleteUser(context, db)
	})
}
