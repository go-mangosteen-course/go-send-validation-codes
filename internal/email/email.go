package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "fangyinghang@foxmail.com")
	m.SetHeader("To", "frankfang1990@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>方应杭</b>!")
	d := gomail.NewDialer(
		viper.GetString("email.smtp.host"),
		viper.GetInt("email.smtp.port"),
		viper.GetString("email.smtp.user"),
		viper.GetString("email.smtp.password"),
	)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
