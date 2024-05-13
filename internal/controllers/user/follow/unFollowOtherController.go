package controllers

import (
	services "BBBingyan/internal/services/user/follow"

	"github.com/labstack/echo/v4"
)

func UnFollowOtherController(c echo.Context) error {
	paramsMap := make(map[string]string)
	paramsMap["unFollowUserId"] = c.QueryParam("un-follow-user-id")
	return services.UnFollowOtherService(paramsMap, c)
}
