package controllers

import (
	services "BBBingyan/internal/services/like"

	"github.com/labstack/echo/v4"
)

func AddNewLikeController(c echo.Context) error {
	toPassageId := c.FormValue("toPassageId")
	paramsMap := map[string]string{
		"toPassageId": toPassageId,
	}
	return services.AddNewLikeService(paramsMap, c)
}
