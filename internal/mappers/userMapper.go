// Package mappers 只负责和数据库交互，不做业务逻辑处理
package mappers

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"

	"github.com/sirupsen/logrus"
)

type UserMapper struct {
}

func (um *UserMapper) AddNewUser(user *dataModels.User) error {
	result := utils.DB.Create(user)
	return result.Error
}

func (um *UserMapper) DeleteUser(user *dataModels.User) error {
	result := utils.DB.Delete(user)
	return result.Error
}

func (um *UserMapper) DeleteUnscopedUser(user *dataModels.User) error {
	result := utils.DB.Unscoped().Delete(user)
	return result.Error
}

func (um *UserMapper) UpdateUser(user *dataModels.User) error {
	result := utils.DB.Save(user)
	return result.Error
}

func (um *UserMapper) GetAllUsers() ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserName(userName string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_name=?", userName)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserEmail(userEmail string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_email=?", userEmail)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserId(userId uint) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "ID=?", userId)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserNickName(userNickName string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_nick_name=?", userNickName)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserIsActive(userIsActive bool) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_is_active=?", userIsActive)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserActivationCode(userActivationCode string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_activation_code=?", userActivationCode)
	return users, result.Error
}

func (um *UserMapper) GetUsersByUserLoginToken(userLoginToken string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Find(&users, "user_login_token=?", userLoginToken)
	return users, result.Error
}

func (um *UserMapper) IfUserExist(userName string) bool {
	var users []*dataModels.User
	_ = utils.DB.Find(&users, "user_name=?", userName)
	if len(users) > 0 {
		if users[0].UserIsActive {
			return true
		}
		err := um.DeleteUnscopedUser(users[0])
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "根据用户名删除未激活用户失败",
			}).Error("根据用户名删除未激活用户失败")
		}
	}
	return false
}

func (um *UserMapper) IfUserExistById(userId uint) bool {
	var users []*dataModels.User
	_ = utils.DB.Find(&users, "ID=?", userId)
	if len(users) > 0 {
		if users[0].UserIsActive {
			return true
		}
		err := um.DeleteUnscopedUser(users[0])
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "根据ID删除未激活用户失败",
			}).Error("根据ID删除未激活用户失败")
		}
	}
	return false
}

func (um *UserMapper) IfUserEmailExist(userEmail string) bool {
	var users []*dataModels.User
	_ = utils.DB.Find(&users, "user_email=?", userEmail)
	if len(users) > 0 {
		if users[0].UserIsActive {
			return true
		}
		err := um.DeleteUnscopedUser(users[0])
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "根据邮箱删除未激活用户失败",
			}).Error("根据邮箱删除未激活用户失败")
		}
	}
	return false
}

func (um *UserMapper) SearchUsersByUserName(userName string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	// 根据粉丝数降序排序
	result := utils.DB.Order("user_fans_count desc").Find(&users, "user_name LIKE ?", "%"+userName+"%")
	return users, result.Error
}

func (um *UserMapper) SearchUsersByUserNickName(userNickName string) ([]*dataModels.User, error) {
	var users []*dataModels.User
	result := utils.DB.Order("user_fans_count desc").Find(&users, "user_nick_name LIKE ?", "%"+userNickName+"%")
	return users, result.Error
}
