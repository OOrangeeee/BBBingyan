package controllers

import (
	services "BBBingyan/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func UserUpdatePasswordController(c echo.Context) error {
	paramMap := make(map[string]string)
	paramMap["newPassword"] = c.FormValue("user-password")
	return services.UserUpdatePassword(paramMap, c)
}
