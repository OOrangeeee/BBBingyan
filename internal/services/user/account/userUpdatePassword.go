package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UserUpdatePassword(paramMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	entryptTool := utils.EncryptionTool{}
	userID := c.Get("userID").(uint)
	users, err := userMapper.GetUsersByUserId(userID)
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
	newPassword := paramMap["newPassword"]
	if newPassword == "" {
		utils.Log.WithField("error_message", "密码不能为空").Error("密码不能为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "密码不能为空",
		})
	}
	if len(newPassword) > 20 || len(newPassword) < 6 {
		utils.Log.WithField("error_message", "密码长度不符合要求").Error("密码长度不符合要求")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "密码长度不符合要求",
		})
	}
	if entryptTool.ComparePassword(user.UserPassword, newPassword) {
		utils.Log.WithField("error_message", "新密码与旧密码相同").Error("新密码与旧密码相同")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "新密码与旧密码相同",
		})
	}
	hashedPassword, err := entryptTool.EncryptPassword(newPassword)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "密码加密失败",
		}).Error("密码加密失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "密码加密失败",
		})
	}
	user.UserPassword = hashedPassword
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户信息失败",
		}).Error("更新用户信息失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户信息失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"succeed_message": "修改密码成功",
	})
}
