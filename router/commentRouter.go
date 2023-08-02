package router

import (
	"LK_blog/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/comments", func(context *gin.Context) {
		controller.GetAllComments(context, db)
	})
	r.GET("/comments/:id", func(context *gin.Context) {
		controller.GetCommentById(context, db)
	})
	r.POST("/comments", func(context *gin.Context) {
		controller.CreateComment(context, db)
	})
	r.PUT("/comments/:id", func(context *gin.Context) {
		controller.UpdateComment(context, db)
	})
	r.DELETE("/comments/:id", func(context *gin.Context) {
		controller.DeleteComment(context, db)
	})
}
