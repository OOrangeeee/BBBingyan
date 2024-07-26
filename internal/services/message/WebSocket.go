package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

// 从http升级为websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
		// return r.Header.Get("Origin") == "https://OOrangeeee.com"
	},
}

var (
	clients     = make(map[uint]*websocket.Conn)
	clientsLock sync.RWMutex
)

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

type message struct {
	UserId  uint   `json:"userId"`
	Message string `json:"message"`
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

	// 获取未发送消息
	messageToSends, err := messageMapper.GetNoSendMessageByToUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取未发送消息失败",
		}).Error("获取未发送消息失败")
	}
	for _, messageToSend := range messageToSends {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(messageToSend.Message)); err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "消息发送失败",
			}).Error("消息发送失败")
			break
		}
		messageToSend.IfSend = true
		err = messageMapper.UpdateMessage(messageToSend)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "更新消息发送状态失败",
			}).Error("更新消息发送状态失败")
			break
		}
	}
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
			IfSend:     false,
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
			// 更新消息发送状态
			newMessage.IfSend = true
			err = messageMapper.UpdateMessage(newMessage)
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "更新消息发送状态失败",
				}).Error("更新消息发送状态失败")
			}
		}
	}
	return nil
}
