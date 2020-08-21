package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

var (
	Option Options
)

type Options struct {
	Enable   bool
	Host     string
	From     string
	Username string
	Password string
	Timeout  int
	Port     int
}

func Send(emailTo, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", Option.From)
	msg.SetHeader("To", emailTo)
	//msg.SetAddressHeader("Cc", config.EmailTestTo, "dorajistyle")
	msg.SetHeader("Subject", subject)
	//msg.SetBody("text/plain", body)
	msg.AddAlternative("text/html", body)
	mailer := gomail.NewDialer(Option.Host, Option.Port, Option.Username, Option.Password)

	err := mailer.DialAndSend(msg)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
