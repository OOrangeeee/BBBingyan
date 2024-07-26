package controllers

import (
	services "BBBingyan/internal/services/email"
	"github.com/labstack/echo/v4"
)

func AdminSendAllEmailsController(c echo.Context) error {
	paramMap := make(map[string]string)
	paramMap["subject"] = c.FormValue("subject")
	paramMap["body"] = c.FormValue("body")
	paramMap["fromNickName"] = c.FormValue("fromNickName")
	return services.SendAllEmails(paramMap, c)
}
