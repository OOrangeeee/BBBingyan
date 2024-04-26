package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "配置文件读取失败",
		}).Panic("Failed to read config file")
	}
}
