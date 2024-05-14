package utils

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type MileTool struct {
}

func (mT *MileTool) SendMail(mailTo []string, subject string, body string, fromNickName string) error {
	userName := viper.GetString("email.emailUserName")
	password := viper.GetString("email.emailPassword")
	host := viper.GetString("email.emailHost")
	port := viper.GetInt("email.emailPort")
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, fromNickName))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "邮件发送失败",
		}).Error("邮件发送失败")
	}
	return err
}
