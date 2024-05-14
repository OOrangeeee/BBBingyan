package controllers

import (
	services "BBBingyan/internal/services/comment"
	"github.com/labstack/echo/v4"
)

func GetCommentByIdController(c echo.Context) error {
	commentId := c.QueryParam("comment-id")
	paramsMap := map[string]string{
		"commentId": commentId,
	}
	return services.GetCommentByIdService(paramsMap, c)
}

func GetCommentsByFromUserIdController(c echo.Context) error {
	fromUserId := c.QueryParam("from-user-id")
	paramsMap := map[string]string{
		"fromUserId": fromUserId,
	}
	return services.GetCommentsByFromUserIdService(paramsMap, c)
}

func GetCommentsByToPassageIdController(c echo.Context) error {
	toPassageId := c.QueryParam("to-passage-id")
	paramsMap := map[string]string{
		"toPassageId": toPassageId,
	}
	return services.GetCommentsByToPassageIdService(paramsMap, c)
}
