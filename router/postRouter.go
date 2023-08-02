package router

import (
	"LK_blog/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/posts", func(context *gin.Context) {
		controller.GetAllPosts(context, db)
	})
	r.GET("/posts/:id", func(context *gin.Context) {
		controller.GetPostById(context, db)
	})
	r.POST("/posts", func(context *gin.Context) {
		controller.CreatePost(context, db)
	})
	r.PUT("/posts/:id", func(context *gin.Context) {
		controller.UpdatePost(context, db)
	})
	r.DELETE("/posts/:id", func(context *gin.Context) {
		controller.DeletePost(context, db)
	})
}
