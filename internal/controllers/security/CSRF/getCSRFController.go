package controllers

import (
	services "BBBingyan/internal/services/security/CSRF"

	"github.com/labstack/echo/v4"
)

func GetCSRFTokenController(c echo.Context) error {
	return services.GetCSRFTokenService(c)
}
