package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
)

type LikeMapper struct {
}

func (lm *LikeMapper) AddNewLike(like *dataModels.Like) error {
	result := utils.DB.Create(like)
	return result.Error
}

func (lm *LikeMapper) GetLikeById(id uint) ([]*dataModels.Like, error) {
	var likes []*dataModels.Like
	result := utils.DB.Find(&likes, "ID=?", id)
	return likes, result.Error
}

func (lm *LikeMapper) GetLikesByFromUserId(fromUserId uint) ([]*dataModels.Like, error) {
	var likes []*dataModels.Like
	result := utils.DB.Find(&likes, "from_user_id=?", fromUserId)
	return likes, result.Error
}

func (lm *LikeMapper) GetLikesByToPassageId(toPassageId uint) ([]*dataModels.Like, error) {
	var likes []*dataModels.Like
	result := utils.DB.Find(&likes, "to_passage_id=?", toPassageId)
	return likes, result.Error
}

func (lm *LikeMapper) DeleteLike(like *dataModels.Like) error {
	result := utils.DB.Delete(like)
	return result.Error
}

func (lm *LikeMapper) GetLikeByFromUserIdAndToPassageId(fromUserId, toPassageId uint) ([]*dataModels.Like, error) {
	var likes []*dataModels.Like
	result := utils.DB.Find(&likes, "from_user_id=? and to_passage_id=?", fromUserId, toPassageId)
	return likes, result.Error
}

func (lm *LikeMapper) IfLikeExist(fromUserId, toPassageId uint) bool {
	var like dataModels.Like
	result := utils.DB.First(&like, "from_user_id=? and to_passage_id=?", fromUserId, toPassageId)
	return result.Error == nil
}
