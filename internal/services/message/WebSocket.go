package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"net/http"
	"sync"
)

type message struct {
	UserId  uint   `json:"userId"`
	Message string `json:"message"`
}

// 从http升级为websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
		// return r.Header.Get("Origin") == "https://OOrangeeee.com"
	},
}

var (
	clients      = make(map[uint]*websocket.Conn)
	clientsLock  sync.RWMutex
	ConnRabbitMQ *amqp.Connection
	raCh         *amqp.Channel
	messageQueue amqp.Queue
)

func InitRabbitMQ() {
	var err error
	ConnRabbitMQ, err = amqp.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "连接RabbitMQ失败",
		}).Panic("连接RabbitMQ失败")
	}
	raCh, err = ConnRabbitMQ.Channel()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "打开RabbitMQ通道失败",
		}).Panic("打开RabbitMQ通道失败")
	}
	messageQueue, err = raCh.QueueDeclare(
		"messageQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "声明RabbitMQ队列失败",
		}).Panic("声明RabbitMQ队列失败")
	}
	startConsumer()
}

// 发送消息到 RabbitMQ 队列
func publishToQueue(body []byte) error {
	err := raCh.Publish(
		"",                // exchange
		messageQueue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "消息发布失败",
		}).Error("消息发布失败")
		return err
	}
	return nil
}

// 启动 RabbitMQ 消费者
func startConsumer() {
	msgs, err := raCh.Consume(
		messageQueue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "设置消费者失败",
		}).Panic("设置消费者失败")
	}

	go func() {
		for msg := range msgs {
			var messageTemp message
			if err := json.Unmarshal(msg.Body, &messageTemp); err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "消息解析失败",
				}).Error("消息解析失败")
				continue
			}

			clientsLock.RLock()
			client, ok := clients[messageTemp.UserId]
			clientsLock.RUnlock()
			if ok {
				if err := client.WriteMessage(websocket.TextMessage, []byte(messageTemp.Message)); err != nil {
					utils.Log.WithFields(logrus.Fields{
						"error":         err,
						"error_message": "WebSocket 消息发送失败",
					}).Error("WebSocket 消息发送失败")
				}
			}
		}
	}()
}

// 注册WebSocket连接
func registerClient(userId uint, conn *websocket.Conn) {
	clientsLock.Lock()
	clients[userId] = conn
	clientsLock.Unlock() // 修改完成后解锁
}

// 移除WebSocket连接
func unregisterClient(userId uint) {
	clientsLock.Lock()
	delete(clients, userId)
	clientsLock.Unlock()
}

func WebsocketHandler(c echo.Context) error {
	messageMapper := &mappers.MessageMapper{}
	// websocket升级
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "websocket升级失败",
		}).Error("websocket升级失败")
		if ws != nil {
			err := ws.Close()
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "websocket关闭失败",
				}).Error("websocket关闭失败")
			}
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "websocket升级失败",
		})
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "websocket关闭失败",
			}).Error("websocket关闭失败")
		}
	}(ws)
	userId := c.Get("userId").(uint)
	defer unregisterClient(userId)
	registerClient(userId, ws)
	// 连接建立后的处理逻辑
	for {
		_, messageByRead, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "websocket读取失败",
				}).Error("websocket读取失败")
			}
			break
		}

		var messageTemp message
		// 解析消息
		if err := json.Unmarshal(messageByRead, &messageTemp); err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"message":       messageByRead,
				"error_message": "消息解析失败",
			}).Error("消息解析失败")
			continue
		}
		newMessage := &dataModels.Message{
			FromUserId: userId,
			ToUserId:   messageTemp.UserId,
			Message:    messageTemp.Message,
		}
		// 存储消息
		err = messageMapper.AddMessage(newMessage)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "存储消息失败",
			}).Error("存储消息失败")
		}
		// 发送消息
		clientsLock.RLock()
		client, ok := clients[messageTemp.UserId]
		clientsLock.RUnlock()
		if ok {
			if err := client.WriteMessage(websocket.TextMessage, []byte(messageTemp.Message)); err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "消息发送失败",
				}).Error("消息发送失败")
				break
			}
		}
		// 发送到RabbitMQ
		err = publishToQueue(messageByRead)
	}
	return nil
}
