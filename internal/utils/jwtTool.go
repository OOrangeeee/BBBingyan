package utils

import (
	"BBBingyan/internal/models/dataModels"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtTool struct{}

func (j *JwtTool) GenerateLoginToken(user *dataModels.User) (string, error) {
	timeTool := TimeTool{}
	timeNow, err := timeTool.GetCurrentTime()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return "", err
	}
	expirationTime := jwt.NewNumericDate(timeNow.Add(time.Hour * 24))
	claims := jwt.MapClaims{
		"UserId":  user.ID,
		"IsAdmin": user.UserIsAdmin,
		"exp":     expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.jwtSecret")))
	if err != nil {
		return "", err
	}
	return t, nil
}
