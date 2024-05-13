package configs

import (
	commentControllers "BBBingyan/internal/controllers/comment"
	passageControllers "BBBingyan/internal/controllers/passage"
	securityCSRFControllers "BBBingyan/internal/controllers/security/CSRF"
	useAccountControllers "BBBingyan/internal/controllers/user/account"
	useFollowControllers "BBBingyan/internal/controllers/user/follow"

	"github.com/labstack/echo/v4"
)

func GetRouterConfig(e *echo.Echo) {
	// 获取用户信息
	e.GET("/users/account", useAccountControllers.GetUserInfoController)
	// 根据用户Id获取用户信息
	e.GET("/users/account/:userId", useAccountControllers.GetUserInfoByIdController)
	// 根据用户名搜索用户
	e.GET("/users/account/search/username", useAccountControllers.SearchUserByUserNameController)
	// 根据用户昵称搜索用户
	e.GET("/users/account/search/nickname", useAccountControllers.SearchUserByUserNickNameController)
	// 获取CSRF Token
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activationCode", useAccountControllers.UserConfirmController)
	// 获得文章标签
	e.GET("/passage/tags", passageControllers.GetPassageTagsController)
	// 根据文章Id获取文章
	e.GET("/passage", passageControllers.GetPassageByIDController)
	// 根据文章标题获取文章
	e.GET("/passages/title", passageControllers.GetPassagesByPassageTitleController)
	// 根据文章作者用户名获取文章
	e.GET("/passages/author/username", passageControllers.GetPassagesByPassageAuthorUserNameController)
	// 根据文章作者昵称获取文章
	e.GET("/passages/author/nickname", passageControllers.GetPassagesByPassageAuthorNickNameController)
	// 根据文章作者Id获取文章
	e.GET("/passages/author/id", passageControllers.GetPassagesByPassageAuthorIdController)
	// 根据文章标签获取文章
	e.GET("/passages/tag", passageControllers.GetPassagesByPassageTagController)
	// 获得最新文章
	e.GET("/passages/last", passageControllers.GetLastPassagesController)
	// 根据Id获得评论
	e.GET("/comment", commentControllers.GetCommentByIdController)
	// 根据评论者Id获得评论
	e.GET("/comment/from-user", commentControllers.GetCommentsByFromUserIdController)
	// 根据文章Id获得评论
	e.GET("/comment/to-passage", commentControllers.GetCommentsByToPassageIdController)
}

func PostRouterConfig(e *echo.Echo) {
	// 注册
	e.POST("/users/account", useAccountControllers.UserRegisterController)
	// 登录
	e.POST("/users/login", useAccountControllers.UserLoginController)
	// 验证用户登录
	e.POST("/users/login/confirm", useAccountControllers.UserLoginConfirmController)
	// 发布文章
	e.POST("/passage", passageControllers.SendPassageController)
	// 发布评论
	e.POST("/comment", commentControllers.SendCommentController)
	// 关注
	e.POST("/users/follow", useFollowControllers.FollowOtherController)
}

func PutRouterConfig(e *echo.Echo) {
	// 修改昵称
	e.PUT("/users/account/nickname", useAccountControllers.UserUpdateNicknameController)
	// 修改密码
	e.PUT("/users/account/password", useAccountControllers.UserUpdatePasswordController)
}

func DeleteRouterConfig(e *echo.Echo) {
}
