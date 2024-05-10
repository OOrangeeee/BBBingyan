package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func UserLoginService(params map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	encryptTool := utils.EncryptionTool{}
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

	userEmailMapper := mappers.UserEmailMapper{}
	if userEmailMapper.IsUserLoginEmailSendInTimeRange(user.UserEmail) {
		utils.Log.WithField("error_message", "邮件发送过于频繁").Error("邮件发送过于频繁")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "邮件发送过于频繁，请等待两分钟再试",
		})
	}

	confirmTokenTool := utils.ConfirmTokenTool{}
	loginToken := confirmTokenTool.GenerateConfirmToken()
	user.UserLoginToken = loginToken
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户Token失败",
		}).Error("更新用户Token失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户Token失败",
		})
	}

	timeTool := utils.TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取当前时间失败",
		})
	}

	mileTool := utils.MileTool{}
	emailBody := viper.GetString("email.emailOfLogin.body")
	emailBody = strings.Replace(emailBody, "{验证码}", loginToken, -1)
	emailBody = strings.Replace(emailBody, "{用户名}", user.UserName, -1)
	emailBody = strings.Replace(emailBody, "{联系电话}", viper.GetString("info.contactPhone"), -1)
	emailBody = strings.Replace(emailBody, "{电子邮件地址}", viper.GetString("info.emailAddress"), -1)
	emailBody = strings.Replace(emailBody, "{官方网站}", viper.GetString("info.webSite"), -1)
	emailBody = strings.Replace(emailBody, "{登录时间}", timeNow.Format("2006-01-02 15:04:05"), -1)
	err = mileTool.SendMail([]string{user.UserEmail}, viper.GetString("email.emailOfLogin.subject"), emailBody, viper.GetString("email.emailFromNickname"))
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "邮件发送失败",
		}).Error("邮件发送失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "邮件发送失败",
		})
	}

	nowUserEmails, err := userEmailMapper.GetUserEmailsByUserEmail(user.UserEmail)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "查询用户邮箱失败",
		})
	}
	// 邮箱不存在在报错
	if len(nowUserEmails) == 0 {
		utils.Log.WithField("error_message", "不存在当前邮箱").Error("不存在当前邮箱")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "不存在当前邮箱",
		})
	}
	nowUserEmail := nowUserEmails[0]
	timeNow, err = timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取当前时间失败",
		})
	}
	nowUserEmail.EmailLastSentOfLogin = timeNow
	err = userEmailMapper.UpdateUserEmail(nowUserEmail)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户邮箱失败",
		}).Error("更新用户邮箱失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户邮箱失败",
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
		"success_message": "登录验证码已发送",
	})
}

func UserLoginConfirmService(paramMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	confirmToken := paramMap["confirmToken"]
	if confirmToken == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "验证码为空",
		}).Error("验证码为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "验证码为空",
		})
	}
	users, err := userMapper.GetUsersByUserLoginToken(confirmToken)
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
	if user.UserLoginToken != confirmToken {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "验证码错误",
		}).Error("验证码错误")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "验证码错误",
		})
	}
	userEmailMapper := mappers.UserEmailMapper{}
	if !userEmailMapper.IsUserLoginEmailSendInTimeRange(user.UserEmail) {
		utils.Log.WithField("error_message", "验证码过期").Error("验证码过期")
	}

	jwtTool := utils.JwtTool{}
	token, err := jwtTool.GenerateLoginToken(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "生成Token失败",
		}).Error("生成Token失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "生成Token失败",
		})
	}
	uuidTool := utils.UUIDTool{}
	user.UserLoginToken = uuidTool.GenerateUUID()
	err = userMapper.UpdateUser(user)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户Token失败",
		}).Error("更新用户Token失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户Token失败",
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
		"success_message": "登录成功",
		"token":           token,
	})
}
