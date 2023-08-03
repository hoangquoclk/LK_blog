package router

import (
	"LK_blog/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/likes", func(context *gin.Context) {
		controller.GetAllLikes(context, db)
	})
	r.POST("/likes", func(context *gin.Context) {
		controller.CreateLike(context, db)
	})
	r.DELETE("/likes/:userId/:postId", func(context *gin.Context) {
		controller.DeleteLike(context, db)
	})
}
