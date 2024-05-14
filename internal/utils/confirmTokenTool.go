package utils

import (
	"github.com/sirupsen/logrus"
	"math/rand"
)

type ConfirmTokenTool struct {
}

func (cT *ConfirmTokenTool) GenerateConfirmToken() string {
	timeTool := TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return ""
	}
	r := rand.New(rand.NewSource(timeNow.UnixNano()))
	confirmToken := make([]rune, 6)
	for i := 0; i < 6; i++ {
		confirmToken[i] = rune(r.Intn(10) + 48)
	}
	return string(confirmToken)
}
