package controllers

import (
	services "BBBingyan/internal/services/user/follow"

	"github.com/labstack/echo/v4"
)

func GetFollowsByFromUserIdController(c echo.Context) error {
	paramsMap := make(map[string]string)
	paramsMap["fromUserId"] = c.QueryParam("from-user-id")
	return services.GetFollowsByFromUserIdService(paramsMap, c)
}

func GetFollowsByToUserIdController(c echo.Context) error {
	paramsMap := make(map[string]string)
	paramsMap["toUserId"] = c.QueryParam("to-user-id")
	return services.GetFollowsByToUserIdService(paramsMap, c)
}
