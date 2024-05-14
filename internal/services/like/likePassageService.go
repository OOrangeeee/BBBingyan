package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AddNewLikeService(paramsMap map[string]string, c echo.Context) error {
	likeMapper := mappers.LikeMapper{}
	passageMapper := mappers.PassageMapper{}
	userMapper := mappers.UserMapper{}
	toPassageId := paramsMap["toPassageId"]
	if toPassageId == "" {
		utils.Log.WithField("error_message", "文章ID为空").Error("文章ID为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章ID为空",
		})
	}
	toPassageIdUint64, err := strconv.ParseUint(toPassageId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文章ID转换失败",
		}).Error("文章ID转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章ID转换失败",
		})
	}
	toPassageIdUint := uint(toPassageIdUint64)
	if !passageMapper.IfPassageExist(toPassageIdUint) {
		utils.Log.WithField("error_message", "文章不存在").Error("文章不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	userId := c.Get("userId").(uint)
	if !userMapper.IfUserExistById(userId) {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	if likeMapper.IfLikeExist(userId, toPassageIdUint) {
		utils.Log.WithField("error_message", "已点赞").Error("已点赞")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "已点赞",
		})
	}
	like := &dataModels.Like{
		FromUserId:  userId,
		ToPassageId: toPassageIdUint,
	}
	err = likeMapper.AddNewLike(like)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "点赞失败",
		}).Error("点赞失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "点赞失败",
		})
	}
	passages, err := passageMapper.GetPassagesByID(toPassageIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取文章失败",
		}).Error("获取文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取文章失败",
		})
	}
	passages[0].PassageBeLikedCount++
	err = passageMapper.UpdatePassage(passages[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新文章失败",
		}).Error("更新文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新文章失败",
		})
	}
	users, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		}).Error("获取用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	users[0].UserLikeCount++
	err = userMapper.UpdateUser(users[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户失败",
		}).Error("更新用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "点赞成功",
	})
}
