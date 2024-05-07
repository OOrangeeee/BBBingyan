package configs

import (
	passageControllers "BBBingyan/internal/controllers/passage"
	securityCSRFControllers "BBBingyan/internal/controllers/security/CSRF"
	useAccountControllers "BBBingyan/internal/controllers/user/account"

	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	// 获取用户信息
	e.GET("/users/account", useAccountControllers.UserGetInfoController)
	// 获取CSRF Token
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activationCode", useAccountControllers.UserConfirmController)
	// 获得文章标签
	e.GET("/passage/tags", passageControllers.GetPassageTagsController)
}

func PostRouterConfig(e *echo.Echo) {
	// 注册
	e.POST("/users/account", useAccountControllers.UserRegisterController)
	// 登录
	e.POST("/users/login", useAccountControllers.UserLoginController)
	// 发布文章
	e.POST("/passage", passageControllers.SendPassageController)
}

func PutRouterConfig(e *echo.Echo) {
	// 修改昵称
	e.PUT("/users/account/nickname", useAccountControllers.UserUpdateNicknameController)
	// 修改密码
	e.PUT("/users/account/password", useAccountControllers.UserUpdatePasswordController)
}

func DeleteRouterConfig(e *echo.Echo) {
}
