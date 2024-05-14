package controllers

import (
	services "BBBingyan/internal/services/passage"

	"github.com/labstack/echo/v4"
)

func SendPassageController(c echo.Context) error {
	passageTitle := c.FormValue("passage-title")
	passageContent := c.FormValue("passage-content")
	passageTag := c.FormValue("passage-tag")
	mapParams := map[string]string{
		"passageTitle":   passageTitle,
		"passageContent": passageContent,
		"passageTag":     passageTag,
	}
	return services.SendPassageService(mapParams, c)
}
