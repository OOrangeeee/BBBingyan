package main

import (
	"BBBingyan/internal/configs"
	"BBBingyan/internal/utils"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	userName := viper.GetString("email.emailUserName")
	password := viper.GetString("email.emailPassword")
	host := viper.GetString("email.emailHost")
	port := viper.GetInt("email.emailPort")
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, "肠镜明眸官方团队"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, password)
	err := d.DialAndSend(m)
	return err
}

func main() {
	utils.InitLog()
	configs.InitViper()
	mailTo := []string{
		"Jin0714@outlook.com",
	}
	subject := viper.GetString("email.emailOfRegister.subject")
	body := viper.GetString("email.emailOfRegister.body")
	username := "张三"                                // 示例用户名
	activationLink := "http://example.com/activate" // 示例激活链接
	contactPhone := viper.GetString("info.contactPhone")
	emailAddress := viper.GetString("info.emailAddress")
	webSite := viper.GetString("info.webSite")
	body = strings.Replace(body, "{用户名}", username, -1)
	body = strings.Replace(body, "{激活链接}", activationLink, -1)
	body = strings.Replace(body, "{联系电话}", contactPhone, -1)
	body = strings.Replace(body, "{电子邮件地址}", emailAddress, -1)
	body = strings.Replace(body, "{官方网站}", webSite, -1)
	println(subject)
	println(body)
	err := SendMail(mailTo, subject, body)
	if err != nil {
		println(err.Error())
	}
	println("邮件发送成功")
}
