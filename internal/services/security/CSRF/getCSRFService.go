package services

import (
	"BBBingyan/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCSRFTokenService(c echo.Context) error {
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if getCSRF == false {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "CSRF Token 获取成功",
	})
}
