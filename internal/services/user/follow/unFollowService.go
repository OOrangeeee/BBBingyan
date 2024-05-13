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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "取消关注成功",
	})
}
