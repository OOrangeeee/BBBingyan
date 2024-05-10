package main

import (
	"BBBingyan/internal/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	timeTool := utils.TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return
	}
	println(timeNow.Format("2006-01-02 15:04:05"))

}
