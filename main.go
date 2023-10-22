package main

import (
	"LK_blog/config"
	"LK_blog/router"
	"github.com/Valgard/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.DatabaseConnection()
	if err != nil {
		panic("failed to connect database")
	}

	dotenv := godotenv.New()

	if err := dotenv.Load(".env"); err != nil {
		panic(err)
	}

	r := gin.Default()

	router.UserRouter(r, db)
	router.CategoryRouter(r, db)
	router.PostRouter(r, db)
	router.CommentRouter(r, db)
	router.LikeRouter(r, db)

	r.Run(":8080")
}
