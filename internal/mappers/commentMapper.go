package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
)

type CommentMapper struct {
}

func (cm *CommentMapper) AddNewComment(comment *dataModels.Comment) error {
	result := utils.DB.Create(comment)
	return result.Error
}

func (cm *CommentMapper) GetCommentById(id uint) ([]*dataModels.Comment, error) {
	var comments []*dataModels.Comment
	result := utils.DB.Find(&comments, "ID=?", id)
	return comments, result.Error
}

func (cm *CommentMapper) GetCommentsByFromUserId(fromUserId uint) ([]*dataModels.Comment, error) {
	var comments []*dataModels.Comment
	result := utils.DB.Find(&comments, "from_user_id=?", fromUserId)
	return comments, result.Error
}

func (cm *CommentMapper) GetCommentsByToPassageId(toPassageId uint) ([]*dataModels.Comment, error) {
	var comments []*dataModels.Comment
	result := utils.DB.Find(&comments, "to_passage_id=?", toPassageId)
	return comments, result.Error
}

func (cm *CommentMapper) DeleteComment(comment *dataModels.Comment) error {
	result := utils.DB.Delete(comment)
	return result.Error
}
