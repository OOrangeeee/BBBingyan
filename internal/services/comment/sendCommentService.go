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

func SendCommentService(paramsMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	passageMapper := mappers.PassageMapper{}
	commentContent := paramsMap["commentContent"]
	toPassageId := paramsMap["toPassageId"]
	if commentContent == "" || toPassageId == "" {
		utils.Log.WithField("error_message", "评论信息不完整").Error("评论信息不完整")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "评论信息不完整",
		})
	}
	userId := c.Get("userId").(uint)
	userTemp, err := userMapper.GetUsersByUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		}).Error("获取用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	if len(userTemp) == 0 {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	userNow := userTemp[0]
	toPassageIdUint64, err := strconv.ParseUint(toPassageId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文章ID转换失败",
		}).Error("文章ID转换失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "文章ID转换失败",
		})
	}
	toPassageIdUint := uint(toPassageIdUint64)
	toPassage, err := passageMapper.GetPassageByID(toPassageIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取文章失败",
		}).Error("获取文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取文章失败",
		})
	}
	if len(toPassage) == 0 {
		utils.Log.WithField("error_message", "文章不存在").Error("文章不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	comment := &dataModels.Comment{
		CommentContent: commentContent,
		FromUserId:     userId,
		ToPassageId:    toPassage[0].ID,
	}
	err = utils.DB.Create(comment).Error
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "评论失败",
		}).Error("评论失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "评论失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "评论成功",
	})
}
