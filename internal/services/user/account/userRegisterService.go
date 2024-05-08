// Package services 业务逻辑代码
package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func RegisterUserService(paramMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	userEmailMapper := mappers.UserEmailMapper{}
	encryptTool := utils.EncryptionTool{}
	userName, userPassword, userEmail, userNickName, userAdminSecret := paramMap["userName"], paramMap["userPassword"], paramMap["userEmail"], paramMap["userNickName"], paramMap["userAdminSecret"]
	if userName == "" || userPassword == "" || userEmail == "" || userNickName == "" {
		utils.Log.WithField("error_message", "用户注册参数不足").Error("用户注册参数不足")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户注册参数不足",
		})
	}
	// 检查管理员密钥
	userIsAdmin := false
	if len(userAdminSecret) > 0 {
		if userAdminSecret != viper.GetString("admin.adminSecret") {
			utils.Log.WithField("error_message", "管理员密钥错误").Error("管理员密钥错误")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "管理员密钥错误",
			})
		}
		userIsAdmin = true
	}
	hashedPassword, err := encryptTool.EncryptPassword(userPassword)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "密码加密失败",
		}).Panic("密码加密失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "密码加密失败",
		})
	}
	userName = strings.TrimSpace(userName)
	// 检查用户名是否合法
	if len(userName) < 6 || len(userName) > 20 {
		utils.Log.WithField("error_message", "用户名长度不合法").Error("用户名长度不合法")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户名长度不合法",
		})
	}
	if !isAlphanumeric(userName) {
		utils.Log.WithField("error_message", "用户名只能包含字母和数字").Error("用户名只能包含字母和数字")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户名只能包含字母和数字",
		})
	}
	// 检查密码是否合法
	if len(userPassword) < 6 || len(userPassword) > 20 {
		utils.Log.WithField("error_message", "密码长度不合法").Error("密码长度不合法")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "密码长度不合法",
		})
	}
	// 检查邮箱是否合法
	if !isValidEmail(userEmail) {
		utils.Log.WithField("error_message", "邮箱格式不合法").Error("邮箱格式不合法")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "邮箱格式不合法",
		})
	}
	// 检查昵称是否合法
	if len(userNickName) < 1 || len(userNickName) > 20 {
		utils.Log.WithField("error_message", "昵称长度不合法").Error("昵称长度不合法")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "昵称长度不合法",
		})
	}
	// 检查用户名是否已存在
	if userMapper.IfUserExist(userName) {
		utils.Log.WithField("error_message", "用户名已存在").Error("用户名已存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户名已存在",
		})
	}
	// 检查邮箱是否已存在
	if userMapper.IfUserEmailExist(userEmail) {
		utils.Log.WithField("error_message", "邮箱已存在").Error("邮箱已存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "邮箱已存在",
		})
	}
	// 注册用户
	newUser := dataModels.User{
		UserName:           userName,
		UserPassword:       hashedPassword,
		UserEmail:          userEmail,
		UserNickName:       userNickName,
		UserIsActive:       false,
		UserActivationCode: getUserActivationCode(),
		UserIsAdmin:        userIsAdmin,
	}
	err = userMapper.AddNewUser(&newUser)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "用户注册失败",
		}).Error("用户注册失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "用户注册失败",
		})
	}
	// 注册用户邮箱
	newUserEmail := dataModels.UserEmail{
		Email: userEmail,
	}
	// 如果邮箱不存在则添加
	if !userEmailMapper.IsExistUserEmail(userEmail) {
		err = userEmailMapper.AddNewUserEmail(&newUserEmail)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "用户邮箱注册失败",
			}).Error("用户邮箱注册失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "用户邮箱注册失败",
			})
		}
	}
	// 发送激活邮件
	// 检查邮件发送时间间隔
	if userEmailMapper.IsUserRegisterEmailSendInTimeRange(userEmail) {
		utils.Log.WithField("error_message", "邮件发送过于频繁").Error("邮件发送过于频繁")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "邮件发送过于频繁，请等待五分钟再试",
		})
	}
	// 发送邮件
	activationURL := viper.GetString("server.host") + "/users/account/activation/" + newUser.UserActivationCode
	mileTool := utils.MileTool{}
	emailBody := viper.GetString("email.emailOfRegister.body")
	emailBody = strings.Replace(emailBody, "{激活链接}", activationURL, -1)
	emailBody = strings.Replace(emailBody, "{用户名}", newUser.UserName, -1)
	emailBody = strings.Replace(emailBody, "{联系电话}", viper.GetString("info.contactPhone"), -1)
	emailBody = strings.Replace(emailBody, "{电子邮件地址}", viper.GetString("info.emailAddress"), -1)
	emailBody = strings.Replace(emailBody, "{官方网站}", viper.GetString("info.webSite"), -1)
	err = mileTool.SendMail([]string{userEmail}, viper.GetString("email.emailOfRegister.subject"), emailBody, viper.GetString("email.emailFromNickname"))
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "邮件发送失败",
		}).Error("邮件发送失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "邮件发送失败",
		})
	}
	// 更新邮件发送时间
	nowUserEmails, err := userEmailMapper.GetUserEmailsByUserEmail(userEmail)
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
	nowUserEmail.EmailLastSentOfRegister = time.Now()
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
	// 返回注册成功信息
	utils.Log.WithFields(logrus.Fields{
		"userName":        userName,
		"userEmail":       userEmail,
		"userNickName":    userNickName,
		"success_message": "用户注册成功，请查收邮件激活账号",
	}).Info("用户注册成功")
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "用户注册成功，请查收邮件激活账号",
	})
}

func isAlphanumeric(s string) bool {
	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	return isAlphanumeric(s)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func getUserActivationCode() string {
	uT := utils.UUIDTool{}
	return uT.GenerateUUID()
}
