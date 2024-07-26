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
