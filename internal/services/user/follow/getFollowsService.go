package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetFollowsByFromUserIdService(paramsMap map[string]string, c echo.Context) error {
	followMapper := mappers.FollowMapper{}
	fromUserIdStr := paramsMap["fromUserId"]
	if fromUserIdStr == "" {
		utils.Log.WithField("error_message", "fromUserId为空").Error("fromUserId为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "fromUserId为空"})
	}
	fromUserIdUint64, err := strconv.ParseUint(fromUserIdStr, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "fromUserId转换失败",
		}).Error("fromUserId转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "fromUserId转换失败"})
	}
	fromUserId := uint(fromUserIdUint64)
	follows, err := followMapper.GetFollowsByFromUserId(fromUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "查询关注失败",
		}).Error("查询关注失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "查询关注失败"})
	}
	userMapper := mappers.UserMapper{}
	var ansFollows []infoModels.Follow
	for _, follow := range follows {
		fromUsers, err := userMapper.GetUsersByUserId(follow.FromUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "查询用户失败",
			}).Error("查询用户失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "查询用户失败"})
		}
		if len(fromUsers) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "用户不存在",
			}).Error("用户不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "用户不存在"})
		}
		fromUser := fromUsers[0]
		toUsers, err := userMapper.GetUsersByUserId(follow.ToUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "查询用户失败",
			}).Error("查询用户失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "查询用户失败"})
		}
		if len(toUsers) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "用户不存在",
			}).Error("用户不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "用户不存在"})
		}
		toUser := toUsers[0]
		fromUserInfo := infoModels.User{
			UserId:           fromUser.ID,
			UserName:         fromUser.UserName,
			UserEmail:        fromUser.UserEmail,
			UserNickName:     fromUser.UserNickName,
			UserFollowCount:  fromUser.UserFollowCount,
			UserFansCount:    fromUser.UserFansCount,
			UserPassageCount: fromUser.UserPassageCount,
			UserLikeCount:    fromUser.UserLikeCount,
			UserIsAdmin:      fromUser.UserIsAdmin,
		}
		toUserInfo := infoModels.User{
			UserId:           toUser.ID,
			UserName:         toUser.UserName,
			UserEmail:        toUser.UserEmail,
			UserNickName:     toUser.UserNickName,
			UserFollowCount:  toUser.UserFollowCount,
			UserFansCount:    toUser.UserFansCount,
			UserPassageCount: toUser.UserPassageCount,
			UserLikeCount:    toUser.UserLikeCount,
			UserIsAdmin:      toUser.UserIsAdmin,
		}
		newFollow := infoModels.Follow{
			ID:       follow.ID,
			FromUser: fromUserInfo,
			ToUser:   toUserInfo,
		}
		ansFollows = append(ansFollows, newFollow)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"follows":         ansFollows,
		"success_message": "查询关注成功"})
}

func GetFollowsByToUserIdService(paramsMap map[string]string, c echo.Context) error {
	followMapper := mappers.FollowMapper{}
	toUserIdStr := paramsMap["toUserId"]
	if toUserIdStr == "" {
		utils.Log.WithField("error_message", "toUserId为空").Error("toUserId为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "toUserId为空"})
	}
	toUserIdUint64, err := strconv.ParseUint(toUserIdStr, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "toUserId转换失败",
		}).Error("toUserId转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "toUserId转换失败"})
	}
	toUserId := uint(toUserIdUint64)
	follows, err := followMapper.GetFollowsByToUserId(toUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "查询关注失败",
		}).Error("查询关注失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "查询关注失败"})
	}
	userMapper := mappers.UserMapper{}
	var ansFollows []infoModels.Follow
	for _, follow := range follows {
		fromUsers, err := userMapper.GetUsersByUserId(follow.FromUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "查询用户失败",
			}).Error("查询用户失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "查询用户失败"})
		}
		if len(fromUsers) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "用户不存在",
			}).Error("用户不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "用户不存在"})
		}
		fromUser := fromUsers[0]
		toUsers, err := userMapper.GetUsersByUserId(follow.ToUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "查询用户失败",
			}).Error("查询用户失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "查询用户失败"})
		}
		if len(toUsers) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "用户不存在",
			}).Error("用户不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "用户不存在"})
		}
		toUser := toUsers[0]
		fromUserInfo := infoModels.User{
			UserId:           fromUser.ID,
			UserName:         fromUser.UserName,
			UserEmail:        fromUser.UserEmail,
			UserNickName:     fromUser.UserNickName,
			UserFollowCount:  fromUser.UserFollowCount,
			UserFansCount:    fromUser.UserFansCount,
			UserPassageCount: fromUser.UserPassageCount,
			UserLikeCount:    fromUser.UserLikeCount,
			UserIsAdmin:      fromUser.UserIsAdmin,
		}
		toUserInfo := infoModels.User{
			UserId:           toUser.ID,
			UserName:         toUser.UserName,
			UserEmail:        toUser.UserEmail,
			UserNickName:     toUser.UserNickName,
			UserFollowCount:  toUser.UserFollowCount,
			UserFansCount:    toUser.UserFansCount,
			UserPassageCount: toUser.UserPassageCount,
			UserLikeCount:    toUser.UserLikeCount,
			UserIsAdmin:      toUser.UserIsAdmin,
		}
		newFollow := infoModels.Follow{
			ID:       follow.ID,
			FromUser: fromUserInfo,
			ToUser:   toUserInfo,
		}
		ansFollows = append(ansFollows, newFollow)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"follows":         ansFollows,
		"success_message": "查询关注成功"})
}
