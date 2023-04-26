package email

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var (
	EMAIL_STMP_HOST = os.Getenv("EMAIL_SMTP_HOST")
	EMAIL_STMP_PORT = os.Getenv("EMAIL_SMTP_PORT")
	EMAIL_USER      = os.Getenv("EMAIL_USER")
	EMAIL_PWD       = os.Getenv("EMAIL_PWD")
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "fangyinghang@foxmail.com")
	m.SetHeader("To", "frankfang1990@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>方应杭</b>!")
	var d *gomail.Dialer
	if port, err := strconv.Atoi(EMAIL_STMP_PORT); err != nil {
		log.Fatalln("EMAIL_STMP_PORT is not a number")
	} else {
		d = gomail.NewDialer(EMAIL_STMP_HOST, port, EMAIL_USER, EMAIL_PWD)
	}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
