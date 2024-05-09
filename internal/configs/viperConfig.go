package configs

import (
	"BBBingyan/internal/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "配置文件读取失败",
		}).Panic("配置文件读取失败")
	}
}
