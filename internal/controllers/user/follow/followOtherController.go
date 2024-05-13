package controllers

import (
	services "BBBingyan/internal/services/user/follow"

	"github.com/labstack/echo/v4"
)

func FollowOtherController(c echo.Context) error {
	paramsMap := make(map[string]string)
	paramsMap["followUserId"] = c.FormValue("follow-user-id")
	return services.UserFollowOtherService(paramsMap, c)
}
