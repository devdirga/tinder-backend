package util

import (
	"fmt"
	"gotinder/config"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Tinder Notif <f88matew@gmail.com>"
const CONFIG_AUTH_EMAIL = "f88matew@gmail.com"

var CONFIG_AUTH_PASSWORD = ""

func SendMail(param map[string]interface{}) error {
	CONFIG_AUTH_PASSWORD = config.GetConf().GoogleSmtpKey
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", param["to"].(string))
	mailer.SetHeader("Subject", param["subject"].(string))
	mailer.SetBody("text/html", fmt.Sprintf(`<p>Click <a href="%s">here</a> to confirm email!</p>`, param["message"].(string)))
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	return err
}

func SendMailResetPassword(param map[string]interface{}) error {
	CONFIG_AUTH_PASSWORD = config.GetConf().GoogleSmtpKey
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", param["to"].(string))
	mailer.SetHeader("Subject", param["subject"].(string))
	mailer.SetBody("text/html", fmt.Sprintf(`<p>Click <a href="%s">here</a> to reset password!</p>`, param["message"].(string)))
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	return err
}
