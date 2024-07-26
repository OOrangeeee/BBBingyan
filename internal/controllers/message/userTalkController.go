package controllers

import (
	services "BBBingyan/internal/services/message"
	"github.com/labstack/echo/v4"
)

func GetWs(c echo.Context) error {
	return services.WebsocketHandler(c)
}
