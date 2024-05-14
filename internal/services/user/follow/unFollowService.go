package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UnFollowOtherService(paramsMap map[string]string, c echo.Context) error {
	unFollowUserIdStr := paramsMap["unFollowUserId"]
	userId := c.Get("userId").(uint)
	if unFollowUserIdStr == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "unFollowUserId 为空",
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "unFollowUserId 为空",
		})
	}
	unFollowUserIdUint64, err := strconv.ParseUint(unFollowUserIdStr, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "unFollowUserId 转换失败",
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "unFollowUserId 转换失败",
		})
	}
	unFollowUserId := uint(unFollowUserIdUint64)
	followMapper := mappers.FollowMapper{}
	follows, err := followMapper.GetFollowsByFromUserIdAndToUserId(userId, unFollowUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取关注失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取关注失败",
		})
	}
	if len(follows) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "未关注该用户",
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "未关注该用户",
		})
	}
	err = followMapper.DeleteFollow(follows[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "取消关注失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "取消关注失败",
		})
	}
	userMapper := mappers.UserMapper{}
	userNow, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	if len(userNow) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "用户不存在",
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	userNow[0].UserFollowCount--
	err = userMapper.UpdateUser(userNow[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户关注数失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户关注数失败",
		})
	}
	followUser, err := userMapper.GetUsersByUserId(unFollowUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取关注用户失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取关注用户失败",
		})
	}
	if len(followUser) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "关注用户不存在",
		})
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "关注用户不存在",
		})
	}
	followUser[0].UserFansCount--
	err = userMapper.UpdateUser(followUser[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新关注用户粉丝数失败",
		})
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新关注用户粉丝数失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "取消关注成功",
	})
}
