package dataModels

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName           string `gorm:"unique"`
	UserPassword       string
	UserEmail          string `gorm:"unique"`
	UserNickName       string
	UserFollowCount    int
	UserFansCount      int
	UserPassageCount   int
	UserLikeCount      int
	UserIsActive       bool   `gorm:"default:false"`
	UserActivationCode string `gorm:"unique"`
	UserIsAdmin        bool   `gorm:"default:false"`
	UserLoginToken     string `gorm:"default:'',unique"`
}
