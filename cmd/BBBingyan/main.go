package main

import (
	"BBBingyan/internal/configs"
	services "BBBingyan/internal/services/message"
	"BBBingyan/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	utils.InitLog()
	configs.InitViper()
	utils.InitDB()
	configs.InitMiddleware(e)
	services.InitRabbitMQ()
	configs.GetRouterConfig(e)
	configs.PostRouterConfig(e)
	configs.PutRouterConfig(e)
	configs.DeleteRouterConfig(e)
	e.Logger.Fatal(e.Start(":714"))
	defer close(utils.Sender.EmailQueue)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "关闭RabbitMQ连接失败",
			}).Panic("关闭RabbitMQ连接失败")
		}
	}(services.ConnRabbitMQ)
}
