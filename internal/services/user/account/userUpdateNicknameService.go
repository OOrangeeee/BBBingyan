package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UserUpdateNicknameService(paramMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userId := c.Get("userId").(uint)
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
	newUserNickName := paramMap["userNickName"]
	if newUserNickName == "" {
		utils.Log.WithField("error_message", "昵称不能为空").Error("昵称不能为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "昵称不能为空",
		})
	}
	if len(newUserNickName) > 20 || len(newUserNickName) < 1 {
		utils.Log.WithField("error_message", "昵称长度不符合要求").Error("昵称长度不符合要求")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "昵称长度不符合要求",
		})
	}
	user.UserNickName = newUserNickName
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户昵称失败",
		}).Error("更新用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户昵称失败",
		})
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "更新用户昵称成功",
	})
}
