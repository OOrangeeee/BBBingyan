package controllers

import (
	services "BBBingyan/internal/services/passage"

	"github.com/labstack/echo/v4"
)

func GetPassageTagsController(c echo.Context) error {
	return services.GetPassageTagsService(c)
}
