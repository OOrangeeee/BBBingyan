package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UserFollowOtherService(paramsMap map[string]string, c echo.Context) error {
	followMapper := mappers.FollowMapper{}
	followUserIdStr := paramsMap["followUserId"]
	if followUserIdStr == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "需要关注的用户ID为空",
		}).Error("需要关注的用户ID为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "需要关注的用户ID为空",
		})
	}
	followUserIdUint64, err := strconv.ParseUint(followUserIdStr, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "需要关注的用户ID转换失败",
		}).Error("需要关注的用户ID转换失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "需要关注的用户ID转换失败",
		})
	}
	followUserId := uint(followUserIdUint64)
	userId := c.Get("userId").(uint)
	if followMapper.IfFollowExist(userId, followUserId) {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "已关注该用户",
		}).Error("已关注该用户")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "已关注该用户",
		})
	}
	userMapper := mappers.UserMapper{}
	followUser, err := userMapper.GetUsersByUserId(followUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取需要关注的用户信息失败",
		}).Error("获取需要关注的用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取需要关注的用户信息失败",
		})
	}
	if len(followUser) == 0 {
		utils.Log.WithField("error_message", "需要关注的用户不存在").Error("需要关注的用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "需要关注的用户不存在",
		})
	}
	followUserNow := followUser[0]
	followUserNow.UserFansCount++
	err = userMapper.UpdateUser(followUserNow)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新需要关注的用户信息失败",
		}).Error("更新需要关注的用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新需要关注的用户信息失败",
		})
	}
	user, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户信息失败",
		}).Error("获取用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户信息失败",
		})
	}
	if len(user) == 0 {
		utils.Log.WithField("error_message", "当前用户不存在").Error("当前用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "当前用户不存在",
		})
	}
	userNow := user[0]
	userNow.UserFollowCount++
	err = userMapper.UpdateUser(userNow)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户关注数失败",
		}).Error("更新用户关注数失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户关注数失败",
		})
	}
	newFollow := &dataModels.Follow{
		FromUserId: userId,
		ToUserId:   followUserId,
	}
	err = followMapper.AddNewFollow(newFollow)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "关注用户失败",
		}).Error("关注用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "关注用户失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "关注用户成功",
	})
}
