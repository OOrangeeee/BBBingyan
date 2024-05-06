package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ConfirmUserService(ActivationCode string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	users, err := userMapper.GetUsersByUserActivationCode(ActivationCode)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "数据库查询失败",
		}).Error("数据库查询失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "数据库查询失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "激活码不存在").Error("激活码不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "激活码不存在",
		})
	}
	user := users[0]
	if user.UserIsActive {
		utils.Log.WithField("error_message", "用户已激活").Error("用户已激活")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户已激活",
		})
	}
	user.UserIsActive = true
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "数据库更新失败",
		}).Error("数据库更新失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "数据库更新失败",
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
		"success_message": "用户激活成功",
	})
}
