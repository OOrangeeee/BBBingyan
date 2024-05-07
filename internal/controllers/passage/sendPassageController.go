package controllers

import (
	services "BBBingyan/internal/services/passage"

	"github.com/labstack/echo/v4"
)

func SendPassageController(c echo.Context) error {
	passageTitle := c.FormValue("passageTitle")
	passageContent := c.FormValue("passageContent")
	passageTag := c.FormValue("passageTag")
	mapParams := map[string]string{
		"passageTitle":   passageTitle,
		"passageContent": passageContent,
		"passageTag":     passageTag,
	}
	return services.SendPassage(mapParams, c)
}
