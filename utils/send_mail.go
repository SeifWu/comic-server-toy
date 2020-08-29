package util

import (
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, body string) error {
	connectionInfo := map[string]string{
		"host":     viper.GetString("formMail.host"),
		"port":     viper.GetString("formMail.port"),
		"user":     viper.GetString("formMail.user"),
		"password": viper.GetString("formMail.password"),
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", viper.GetString("formMail.user"))
	mail.SetHeader("To", mailTo...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	port, _ := strconv.Atoi(connectionInfo["port"])

	dialer := gomail.NewDialer(connectionInfo["host"], port, connectionInfo["user"], connectionInfo["password"])

	err := dialer.DialAndSend(mail)
	return err
}
