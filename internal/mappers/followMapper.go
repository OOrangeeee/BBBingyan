package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
)

type FollowMapper struct {
}

func (fm *FollowMapper) AddNewFollow(follow *dataModels.Follow) error {
	result := utils.DB.Create(follow)
	return result.Error
}

func (fm *FollowMapper) DeleteFollow(follow *dataModels.Follow) error {
	result := utils.DB.Delete(follow)
	return result.Error
}

func (fm *FollowMapper) UpdateFollow(follow *dataModels.Follow) error {
	result := utils.DB.Save(follow)
	return result.Error
}

func (fm *FollowMapper) GetAllFollows() ([]*dataModels.Follow, error) {
	var follows []*dataModels.Follow
	result := utils.DB.Find(&follows)
	return follows, result.Error
}

func (fm *FollowMapper) GetFollowsByFollowId(followId uint) ([]*dataModels.Follow, error) {
	var follows []*dataModels.Follow
	result := utils.DB.Find(&follows, "ID=?", followId)
	return follows, result.Error
}

func (fm *FollowMapper) GetFollowsByFromUserId(fromUserId uint) ([]*dataModels.Follow, error) {
	var follows []*dataModels.Follow
	result := utils.DB.Find(&follows, "from_user_id=?", fromUserId)
	return follows, result.Error
}

func (fm *FollowMapper) GetFollowsByToUserId(toUserId uint) ([]*dataModels.Follow, error) {
	var follows []*dataModels.Follow
	result := utils.DB.Find(&follows, "to_user_id=?", toUserId)
	return follows, result.Error
}
