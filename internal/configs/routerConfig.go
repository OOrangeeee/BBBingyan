package configs

import (
	commentControllers "BBBingyan/internal/controllers/comment"
	emailControllers "BBBingyan/internal/controllers/email"
	likeControllers "BBBingyan/internal/controllers/like"
	wsControllers "BBBingyan/internal/controllers/message"
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
	e.GET("/users/account/:user-id", useAccountControllers.GetUserInfoByIdController)
	// 根据用户名搜索用户
	e.GET("/users/account/search/username", useAccountControllers.SearchUserByUserNameController)
	// 根据用户昵称搜索用户
	e.GET("/users/account/search/nickname", useAccountControllers.SearchUserByUserNickNameController)
	// 获取CSRF Token
	e.GET("/csrf-token", securityCSRFControllers.GetCSRFTokenController)
	// 激活
	e.GET("/users/account/activation/:activation-code", useAccountControllers.UserConfirmController)
	// 获得文章标签
	e.GET("/passage/tags", passageControllers.GetPassageTagsController)
	// 根据文章Id获取文章
	e.GET("/passage/:passage-id", passageControllers.GetPassageByIDController)
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
	// 获得最热文章
	e.GET("/passages/hottest", passageControllers.GetHottestPassagesController)
	// 搜索文章
	e.GET("/passages/search", passageControllers.SearchPassagesController)
	// 根据Id获得评论
	e.GET("/comment", commentControllers.GetCommentByIdController)
	// 根据评论者Id获得评论
	e.GET("/comment/from-user", commentControllers.GetCommentsByFromUserIdController)
	// 根据文章Id获得评论
	e.GET("/comment/to-passage", commentControllers.GetCommentsByToPassageIdController)
	// 根据fromUserId获取关注列表
	e.GET("/users/follow/fromuser", useFollowControllers.GetFollowsByFromUserIdController)
	// 根据toUserId获取粉丝列表
	e.GET("/users/follow/touser", useFollowControllers.GetFollowsByToUserIdController)
	// 获取用户点赞列表
	e.GET("/like/:from-user-id", likeControllers.GetLikesByFromUserIdController)
	// ws
	e.GET("/ws", wsControllers.GetWs)
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
	// 点赞
	e.POST("/like", likeControllers.AddNewLikeController)
	// 发送邮件
	e.POST("/email/users", emailControllers.AdminSendAllEmailsController)
}

func PutRouterConfig(e *echo.Echo) {
	// 修改昵称
	e.PUT("/users/account/nickname", useAccountControllers.UserUpdateNicknameController)
	// 修改密码
	e.PUT("/users/account/password", useAccountControllers.UserUpdatePasswordController)
}

func DeleteRouterConfig(e *echo.Echo) {
	// 取消关注
	e.DELETE("/users/follow", useFollowControllers.UnFollowOtherController)
	// 删除文章
	e.DELETE("/passage", passageControllers.DeletePassageController)
	// 取消点赞
	e.DELETE("/like", likeControllers.UnLikeController)
}
