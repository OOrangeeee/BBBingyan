package controllers

import (
	services "BBBingyan/internal/services/passage"

	"github.com/labstack/echo/v4"
)

func GetPassageByIDController(c echo.Context) error {
	id := c.Param("id")
	mapParams := map[string]string{
		"id": id,
	}
	return services.GetPassageByIDService(mapParams, c)
}

func GetPassagesByPassageTitleController(c echo.Context) error {
	title := c.QueryParam("passage-title")
	mapParams := map[string]string{
		"passageTitle": title,
	}
	return services.GetPassagesByPassageTitleService(mapParams, c)
}

func GetPassagesByPassageAuthorUserNameController(c echo.Context) error {
	passageAuthorUserName := c.QueryParam("passage-author-username")
	mapParams := map[string]string{
		"passageAuthorUserName": passageAuthorUserName,
	}
	return services.GetPassagesByPassageAuthorUserNameService(mapParams, c)
}

func GetPassagesByPassageAuthorNickNameController(c echo.Context) error {
	passageAuthorNickName := c.QueryParam("passage-author-nickname")
	mapParams := map[string]string{
		"passageAuthorNickName": passageAuthorNickName,
	}
	return services.GetPassagesByPassageAuthorNickNameService(mapParams, c)
}

func GetPassagesByPassageAuthorIdController(c echo.Context) error {
	return services.GetPassagesByPassageAuthorIdService(c)
}

func GetPassagesByPassageTagController(c echo.Context) error {
	tag := c.QueryParam("passage-tag")
	mapParams := map[string]string{
		"passageTag": tag,
	}
	return services.GetPassagesByPassageTagService(mapParams, c)
}

func GetLastPassagesController(c echo.Context) error {
	return services.GetLastPassagesService(c)
}
