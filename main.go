package main

import (
	"LK_blog/config"
	"LK_blog/router"
	"github.com/gin-gonic/gin"
	"github.com/wneessen/go-mail"
	"log"
)

func main() {
	m := mail.NewMsg()
	if err := m.From("lkblog1999@gmail.com"); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To("hoangquoclk003@gmail.com"); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("Hello cau!")
	m.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")
	_, err := mail.NewClient("smtp.gmail.com", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("lkblog1999@gmail.com"), mail.WithPassword("uroxyzihhzyfxluh"))
	if err != nil {
		log.Fatalf("failed to create mail client: %s", err)
	}
	//TODO tao ham khi tao moi 1 user
	//if err := c.DialAndSend(m); err != nil {
	//	log.Fatalf("failed to send mail: %s", err)
	//}

	db, err := config.DatabaseConnection()

	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	router.UserRouter(r, db)
	router.CategoryRouter(r, db)
	router.PostRouter(r, db)
	router.CommentRouter(r, db)
	router.LikeRouter(r, db)

	r.Run(":8080")
}
