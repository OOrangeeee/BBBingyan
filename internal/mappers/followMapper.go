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

func (fm *FollowMapper) GetFollowsByFromUserIdAndToUserId(fromUserId, toUserId uint) ([]*dataModels.Follow, error) {
	var follows []*dataModels.Follow
	result := utils.DB.Find(&follows, "from_user_id=? AND to_user_id=?", fromUserId, toUserId)
	return follows, result.Error
}

func (fm *FollowMapper) IfFollowExist(fromUserId, toUserId uint) bool {
	follows, err := fm.GetFollowsByFromUserIdAndToUserId(fromUserId, toUserId)
	if err != nil {
		utils.Log.WithField("error_message", err).Error("获取关注失败")
		return false
	}
	if len(follows) == 0 {
		return false
	}
	return true
}
