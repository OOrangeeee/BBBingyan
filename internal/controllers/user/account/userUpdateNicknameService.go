package controllers

import (
	services "BBBingyan/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func UserUpdateNicknameController(c echo.Context) error {
	paramMap := make(map[string]string)
	paramMap["userNickName"] = c.FormValue("user-nickname")
	return services.UserUpdateNicknameService(paramMap, c)
}
