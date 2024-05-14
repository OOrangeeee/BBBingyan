package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCSRFTokenService(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "CSRF Token 获取成功",
	})
}
