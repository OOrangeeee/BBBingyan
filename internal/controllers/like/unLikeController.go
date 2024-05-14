package controllers

import (
	"github.com/labstack/echo/v4"

	services "BBBingyan/internal/services/like"
)

func UnLikeController(c echo.Context) error {
	paramsMap := make(map[string]string)
	passageId := c.QueryParam("passage-id")
	paramsMap["passageId"] = passageId
	return services.UnLikeService(paramsMap, c)
}
