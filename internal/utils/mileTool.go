package utils

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type MileTool struct {
}

/*func (mT *MileTool) SendMail(mailTo []string, subject string, body string, fromNickName string) error {
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
}*/

func (mT *MileTool) SendMailForOne(mailTo string, subject string, body string, fromNickName string) error {
	userName := viper.GetString("email.emailUserName")
	password := viper.GetString("email.emailPassword")
	host := viper.GetString("email.emailHost")
	port := viper.GetInt("email.emailPort")
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, fromNickName))
	m.SetHeader("To", mailTo)
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

// 手写异步邮件发送池

type EmailTask struct {
	Recipient    string
	Subject      string
	Body         string
	FromNickName string
}

// EmailSender 定义了发送邮件需要的信息和方法
type EmailSender struct {
	EmailQueue     chan EmailTask
	WorkerPoolSize int
	MailTool       MileTool
}

var Sender *EmailSender

func init() {
	Sender = NewEmailSender(100, MileTool{})
	Sender.Start()
}

// NewEmailSender 初始化邮件发送器
func NewEmailSender(poolSize int, mailTool MileTool) *EmailSender {
	return &EmailSender{
		EmailQueue:     make(chan EmailTask, 1000), // 调整缓冲大小以适应负载
		WorkerPoolSize: poolSize,
		MailTool:       mailTool,
	}
}

// Start 启动邮件发送协程池
func (sender *EmailSender) Start() {
	for i := 0; i < sender.WorkerPoolSize; i++ {
		go sender.worker()
	}
}

// worker 单个工作协程的处理逻辑
func (sender *EmailSender) worker() {
	for task := range sender.EmailQueue {
		err := sender.MailTool.SendMailForOne(task.Recipient, task.Subject, task.Body, task.FromNickName)
		if err != nil {
			Log.WithFields(logrus.Fields{
				"recipient":     task.Recipient,
				"email_subject": task.Subject,
				"error":         err,
				"error_message": "邮件发送失败",
			}).Error("邮件发送时出错")
		}
	}
}
