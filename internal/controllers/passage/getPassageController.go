package controllers

import (
	services "BBBingyan/internal/services/passage"

	"github.com/labstack/echo/v4"
)

func GetPassageByIDController(c echo.Context) error {
	id := c.Param("passage-id")
	mapParams := map[string]string{
		"id": id,
	}
	return services.GetPassageByIDService(mapParams, c)
}

func GetPassagesByPassageTitleController(c echo.Context) error {
	title := c.QueryParam("passage-title")
	pageSize := c.QueryParam("page-size")
	mapParams := map[string]string{
		"passageTitle": title,
		"pageSize":     pageSize,
	}
	return services.GetPassagesByPassageTitleService(mapParams, c)
}

func GetPassagesByPassageAuthorUserNameController(c echo.Context) error {
	passageAuthorUserName := c.QueryParam("passage-author-username")
	pageSize := c.QueryParam("page-size")
	mapParams := map[string]string{
		"passageAuthorUserName": passageAuthorUserName,
		"pageSize":              pageSize,
	}
	return services.GetPassagesByPassageAuthorUserNameService(mapParams, c)
}

func GetPassagesByPassageAuthorNickNameController(c echo.Context) error {
	passageAuthorNickName := c.QueryParam("passage-author-nickname")
	pageSize := c.QueryParam("page-size")
	mapParams := map[string]string{
		"passageAuthorNickName": passageAuthorNickName,
		"pageSize":              pageSize,
	}
	return services.GetPassagesByPassageAuthorNickNameService(mapParams, c)
}

func GetPassagesByPassageAuthorIdController(c echo.Context) error {
	pageSize := c.QueryParam("page-size")
	mapParams := map[string]string{
		"pageSize": pageSize,
	}
	return services.GetPassagesByPassageAuthorIdService(mapParams, c)
}

func GetPassagesByPassageTagController(c echo.Context) error {
	pageSize := c.QueryParam("page-size")
	tag := c.QueryParam("passage-tag")
	mapParams := map[string]string{
		"passageTag": tag,
		"pageSize":   pageSize,
	}
	return services.GetPassagesByPassageTagService(mapParams, c)
}

func GetLastPassagesController(c echo.Context) error {
	return services.GetLastPassagesService(c)
}

func SearchPassagesController(c echo.Context) error {
	title := c.QueryParam("passage-title")
	pageSize := c.QueryParam("page-size")
	mapParams := map[string]string{
		"searchTitle": title,
		"pageSize":    pageSize,
	}
	return services.SearchPassagesService(mapParams, c)
}
