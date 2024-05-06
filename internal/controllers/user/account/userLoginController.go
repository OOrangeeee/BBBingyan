package controllers

import (
	services "BBBingyan/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func UserLoginController(c echo.Context) error {
	paramMap := make(map[string]string)
	paramMap["userName"] = c.FormValue("user-name")
	paramMap["password"] = c.FormValue("user-password")
	return services.UserLoginService(paramMap, c)
}
