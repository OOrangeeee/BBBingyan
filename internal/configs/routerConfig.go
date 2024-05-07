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
	// 根据文章Id获取文章
	e.GET("/passage/:id", passageControllers.GetPassageByIDController)
	// 根据文章标题获取文章
	e.GET("/passages/title/:passage-title", passageControllers.GetPassagesByPassageTitleController)
	// 根据文章作者用户名获取文章
	e.GET("/passages/author/username/:passage-author-username", passageControllers.GetPassagesByPassageAuthorUserNameController)
	// 根据文章作者昵称获取文章
	e.GET("/passages/author/nickname/:passage-author-nickname", passageControllers.GetPassagesByPassageAuthorNickNameController)
	// 根据文章作者Id获取文章
	e.GET("/passages/author/id", passageControllers.GetPassagesByPassageAuthorIdController)
	// 根据文章标签获取文章
	e.GET("/passages/tag/:passage-tag", passageControllers.GetPassagesByPassageTagController)
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
