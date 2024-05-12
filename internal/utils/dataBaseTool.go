package utils

import (
	"BBBingyan/internal/models/dataModels"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dataBaseUserName string
var dataBasePassword string
var dataBaseIp string
var dataBasePort string
var dataBaseName string

func InitDB() {
	dataBaseName = viper.GetString("database.dataBaseName")
	dataBaseUserName = viper.GetString("database.dataBaseUserName")
	dataBasePassword = viper.GetString("database.dataBasePassword")
	dataBaseIp = viper.GetString("database.dataBaseIp")
	dataBasePort = viper.GetString("database.dataBasePort")
	var err error
	maxRetries := 100
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open("postgres://"+dataBaseUserName+":"+dataBasePassword+"@"+dataBaseIp+":"+dataBasePort+"/"+dataBaseName), &gorm.Config{})
		if err == nil {
			break
		}
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "链接数据库失败，正在重试...",
		}).Error("链接数据库失败，正在重试...")
		time.Sleep(time.Second * 5)
	}
	if DB == nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "链接数据库失败, 重试次数超过最大重试次数",
		}).Panic("链接数据库失败, 重试次数超过最大重试次数")
	}
	DB.Model(&dataModels.User{})
	err = DB.AutoMigrate(&dataModels.User{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建用户表失败",
		}).Panic("创建用户表失败")
	}
	DB.Model(&dataModels.UserEmail{})
	err = DB.AutoMigrate(&dataModels.UserEmail{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建用户邮箱表失败",
		}).Panic("创建用户邮箱表失败")
	}
	DB.Model(&dataModels.Passage{})
	err = DB.AutoMigrate(&dataModels.Passage{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建文章表失败",
		}).Panic("创建文章表失败")
	}
	DB.Model(&dataModels.Comment{})
	err = DB.AutoMigrate(&dataModels.Comment{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建评论表失败",
		}).Panic("创建评论表失败")
	}
	DB.Model(&dataModels.Follow{})
	err = DB.AutoMigrate(&dataModels.Follow{})
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "创建关注表失败",
		}).Panic("创建关注表失败")
	}
}
