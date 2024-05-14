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

func GetUserInfoService(c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userId := c.Get("userId").(uint)
	isAdmin := c.Get("isAdmin").(bool)
	users, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户信息失败",
		}).Error("获取用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户信息失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	user := users[0]
	userInfo := infoModels.User{
		UserId:           user.ID,
		UserName:         user.UserName,
		UserEmail:        user.UserEmail,
		UserNickName:     user.UserNickName,
		UserFollowCount:  user.UserFollowCount,
		UserFansCount:    user.UserFansCount,
		UserPassageCount: user.UserPassageCount,
		UserLikeCount:    user.UserLikeCount,
		UserIsAdmin:      isAdmin,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户信息成功",
		"userInfo":        userInfo,
	})
}

func GetUserInfoByIdService(paramsMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userIdStr := paramsMap["userId"]
	userIdUint64, err := strconv.ParseUint(userIdStr, 10, 64)
	userId := uint(userIdUint64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "用户ID转换失败",
		}).Error("用户ID转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户ID转换失败",
		})
	}
	users, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户信息失败",
		}).Error("获取用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户信息失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	user := users[0]
	userInfo := infoModels.User{
		UserId:           user.ID,
		UserName:         user.UserName,
		UserEmail:        user.UserEmail,
		UserNickName:     user.UserNickName,
		UserFollowCount:  user.UserFollowCount,
		UserFansCount:    user.UserFansCount,
		UserPassageCount: user.UserPassageCount,
		UserLikeCount:    user.UserLikeCount,
		UserIsAdmin:      user.UserIsAdmin,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取用户信息成功",
		"userInfo":        userInfo,
	})
}

func SearchUserByUserNameService(paramsMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userName := paramsMap["userName"]
	users, err := userMapper.SearchUsersByUserName(userName)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "搜索用户失败",
		}).Error("搜索用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "搜索用户失败",
		})
	}
	var userInfos []infoModels.User
	for _, user := range users {
		userInfo := infoModels.User{
			UserId:           user.ID,
			UserName:         user.UserName,
			UserEmail:        user.UserEmail,
			UserNickName:     user.UserNickName,
			UserFollowCount:  user.UserFollowCount,
			UserFansCount:    user.UserFansCount,
			UserPassageCount: user.UserPassageCount,
			UserLikeCount:    user.UserLikeCount,
			UserIsAdmin:      user.UserIsAdmin,
		}
		userInfos = append(userInfos, userInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "搜索用户成功",
		"userInfos":       userInfos,
	})
}

func SearchUserByUserNicknameService(paramsMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userNickName := paramsMap["userNickName"]
	users, err := userMapper.SearchUsersByUserNickName(userNickName)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "搜索用户失败",
		}).Error("搜索用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "搜索用户失败",
		})
	}
	var userInfos []infoModels.User
	for _, user := range users {
		userInfo := infoModels.User{
			UserId:           user.ID,
			UserName:         user.UserName,
			UserEmail:        user.UserEmail,
			UserNickName:     user.UserNickName,
			UserFollowCount:  user.UserFollowCount,
			UserFansCount:    user.UserFansCount,
			UserPassageCount: user.UserPassageCount,
			UserLikeCount:    user.UserLikeCount,
			UserIsAdmin:      user.UserIsAdmin,
		}
		userInfos = append(userInfos, userInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "搜索用户成功",
		"userInfos":       userInfos,
	})
}
