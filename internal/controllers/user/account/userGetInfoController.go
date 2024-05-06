package controllers

import (
	services "BBBingyan/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func UserGetInfoController(c echo.Context) error {
	return services.GetUserInfoService(c)
}
