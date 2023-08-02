package router

import (
	"LK_blog/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/categories", func(context *gin.Context) {
		controller.GetAllCategories(context, db)
	})
	r.GET("/categories/:id", func(context *gin.Context) {
		controller.GetCategoryById(context, db)
	})
	r.POST("/categories", func(context *gin.Context) {
		controller.CreateCategory(context, db)
	})
	r.PUT("/categories/:id", func(context *gin.Context) {
		controller.UpdateCategory(context, db)
	})
	r.DELETE("/categories/:id", func(context *gin.Context) {
		controller.DeleteCategory(context, db)
	})
}
