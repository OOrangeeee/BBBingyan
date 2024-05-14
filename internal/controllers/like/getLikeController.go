package controllers

import (
	services "BBBingyan/internal/services/like"

	"github.com/labstack/echo/v4"
)

func GetLikesByFromUserIdController(c echo.Context) error {
	paramsMap := make(map[string]string)
	fromUserId := c.Param("from-user-id")
	paramsMap["fromUserId"] = fromUserId
	return services.GetUserLikesByUserIdService(paramsMap, c)
}
