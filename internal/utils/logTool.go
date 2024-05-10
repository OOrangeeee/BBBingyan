package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

var Log *logrus.Logger

func InitLog() {
	Log = logrus.New()
	Log.Formatter = &logrus.JSONFormatter{}
	Log.SetReportCaller(true)
	timeTool := TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
	}
	logFileLocation := filepath.Join("./logs", timeNow.Format("2006-01-02")+".log")
	Log.Out = &lumberjack.Logger{
		Filename:   logFileLocation,
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     30,
	}
}
