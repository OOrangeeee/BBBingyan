package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
)

type MessageMapper struct{}

func (m *MessageMapper) AddMessage(message *dataModels.Message) error {
	result := utils.DB.Create(message)
	return result.Error
}

func (m *MessageMapper) GetNoSendMessageByToUserId(toUserId uint) ([]*dataModels.Message, error) {
	var messages []*dataModels.Message
	result := utils.DB.Find(&messages, "to_user_id=? and if_send=?", toUserId, false)
	return messages, result.Error
}

func (m *MessageMapper) UpdateMessage(message *dataModels.Message) error {
	result := utils.DB.Save(message)
	return result.Error
}
