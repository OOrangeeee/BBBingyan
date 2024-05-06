package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UserLoginService(params map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	encryptTool := utils.EncryptionTool{}
	jwtTool := utils.JwtTool{}
	userName := params["userName"]
	password := params["password"]
	if userName == "" || password == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "用户名或密码为空",
		}).Error("用户名或密码为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户名或密码为空",
		})
	}
	users, err := userMapper.GetUsersByUserName(userName)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "查询用户失败",
		}).Error("查询用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "查询用户失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "用户不存在",
		}).Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	user := users[0]
	if !encryptTool.ComparePassword(user.UserPassword, password) {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "密码错误",
		}).Error("密码错误")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "密码错误",
		})
	}
	t, err := jwtTool.GenerateLoginToken(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "生成Token失败",
		}).Error("生成Token失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "生成Token失败",
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
		"token":           t,
		"success_message": "登录成功",
	})
}
