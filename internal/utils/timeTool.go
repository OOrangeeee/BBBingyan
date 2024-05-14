package utils

import (
	"github.com/sirupsen/logrus"
	"time"
)

type TimeTool struct {
}

func (t *TimeTool) GetCurrentTime() (time.Time, error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "时区加载失败",
		}).Error("获取当前时间失败")
		return time.Time{}, err
	}
	now := time.Now().In(location)
	return now, nil
}
