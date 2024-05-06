package utils

import (
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/models/infoModels"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtTool struct{}

func (j *JwtTool) GenerateLoginToken(user *dataModels.User) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	claims := infoModels.JwtCustomClaim{
		UserId:           user.ID,
		IsAdmin:          user.UserIsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: expirationTime},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt.jwtSecret")))
	if err != nil {
		return "", err
	}
	return t, nil
}
