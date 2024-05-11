package controllers

import (
	services "BBBingyan/internal/services/user/account"

	"github.com/labstack/echo/v4"
)

func GetUserInfoController(c echo.Context) error {
	return services.GetUserInfoService(c)
}

func GetUserInfoByIdController(c echo.Context) error {
	userId := c.Param("userId")
	paramsMap := map[string]string{
		"userId": userId,
	}
	return services.GetUserInfoByIdService(paramsMap, c)
}

func SearchUserByUserNameController(c echo.Context) error {
	userName := c.QueryParam("userName")
	paramsMap := map[string]string{
		"userName": userName,
	}
	return services.SearchUserByUserNameService(paramsMap, c)
}

func SearchUserByUserNickNameController(c echo.Context) error {
	userNickName := c.QueryParam("userNickName")
	paramsMap := map[string]string{
		"userNickName": userNickName,
	}
	return services.SearchUserByUserNicknameService(paramsMap, c)
}
