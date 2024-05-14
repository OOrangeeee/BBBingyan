package controllers

import (
	services "BBBingyan/internal/services/passage"
	"github.com/labstack/echo/v4"
)

func DeletePassageController(c echo.Context) error {
	passageId := c.FormValue("passage-id")
	paramsMap := map[string]string{
		"passageId": passageId,
	}
	return services.DeletePassageService(paramsMap, c)
}
