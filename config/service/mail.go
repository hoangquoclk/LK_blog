package service

import (
	"LK_blog/model"
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

type MailInfo struct {
	FirstName string
	LastName  string
}

func SendMail(i model.User) {
	t := template.New("template.html")

	var err error
	t, err = t.ParseFiles("config/service/template.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer

	mailInfo := MailInfo{FirstName: i.FirstName, LastName: i.LastName}

	if err := t.Execute(&tpl, mailInfo); err != nil {
		log.Println(err)
	}

	result := tpl.String()
	m := gomail.NewMessage()

	m.SetHeader("From", "lkblog1999@gmail.com")
	m.SetHeader("To", "hoangquoclk003@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", result)
	d := gomail.NewDialer("smtp.gmail.com", 587, "lkblog1999@gmail.com", "tyypkiafjgwveszd")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
