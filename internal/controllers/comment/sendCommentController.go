package controllers

import (
	services "BBBingyan/internal/services/comment"
	"github.com/labstack/echo/v4"
)

func SendCommentController(c echo.Context) error {
	commentContent := c.FormValue("comment-content")
	toPassageId := c.FormValue("to-passage-id")
	mapParams := map[string]string{
		"commentContent": commentContent,
		"toPassageId":    toPassageId,
	}
	return services.SendCommentService(mapParams, c)
}
