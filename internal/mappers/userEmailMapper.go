package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserEmailMapper struct {
}

func (uem *UserEmailMapper) AddNewUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Create(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) DeleteUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Delete(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) UpdateUserEmail(userEmail *dataModels.UserEmail) error {
	result := utils.DB.Save(userEmail)
	return result.Error
}

func (uem *UserEmailMapper) GetAllUserEmails() ([]*dataModels.UserEmail, error) {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails)
	return userEmails, result.Error
}

func (uem *UserEmailMapper) GetUserEmailsByUserEmail(userEmail string) ([]*dataModels.UserEmail, error) {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails, "email=?", userEmail)
	return userEmails, result.Error
}

func (uem *UserEmailMapper) IsUserRegisterEmailSendInTimeRange(email string) bool {
	timeRange := viper.GetInt("email.emailOfRegister.timeRange")
	timeTool := utils.TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return false
	}
	beforeTime := timeNow.Add(-time.Duration(timeRange) * time.Minute)
	var userEmail *dataModels.UserEmail
	result := utils.DB.First(&userEmail, "email=?", email)
	if result.Error != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         result.Error,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
	}
	if userEmail.EmailLastSentOfRegister.After(beforeTime) {
		return true
	}
	return false
}

func (uem *UserEmailMapper) IsUserLoginEmailSendInTimeRange(email string) bool {
	timeRange := viper.GetInt("email.emailOfLogin.timeRange")
	timeTool := utils.TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return false
	}
	beforeTime := timeNow.Add(-time.Duration(timeRange) * time.Minute)
	var userEmail *dataModels.UserEmail
	result := utils.DB.First(&userEmail, "email=?", email)
	if result.Error != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         result.Error,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
	}
	if userEmail.EmailLastSentOfLogin.After(beforeTime) {
		return true
	}
	return false
}

func (uem *UserEmailMapper) IsExistUserEmail(email string) bool {
	var userEmails []*dataModels.UserEmail
	result := utils.DB.Find(&userEmails, "email=?", email)
	if result.Error != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         result.Error,
			"error_message": "查询用户邮箱失败",
		}).Error("查询用户邮箱失败")
	}
	if len(userEmails) == 0 {
		return false
	}
	return true
}
