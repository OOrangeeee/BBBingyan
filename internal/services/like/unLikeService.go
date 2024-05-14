package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func UnLikeService(paramsMap map[string]string, c echo.Context) error {
	passageId := paramsMap["passageId"]
	if passageId == "" {
		utils.Log.WithField("error_message", "文章ID为空").Error("文章ID为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章ID为空",
		})
	}
	passageIdUint64, err := strconv.ParseUint(passageId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "文章ID转换失败",
		}).Error("文章ID转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章ID转换失败",
		})
	}
	passageIdUint := uint(passageIdUint64)
	userId := c.Get("userId").(uint)
	likeMapper := mappers.LikeMapper{}
	if !likeMapper.IfLikeExist(userId, passageIdUint) {
		utils.Log.WithField("error_message", "未点赞").Error("未点赞")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "未点赞",
		})
	}
	likes, err := likeMapper.GetLikeByFromUserIdAndToPassageId(userId, passageIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取点赞失败",
		}).Error("获取点赞失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取点赞失败",
		})
	}
	if len(likes) == 0 {
		utils.Log.WithField("error_message", "点赞不存在").Error("点赞不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "点赞不存在",
		})
	}
	like := likes[0]
	err = likeMapper.DeleteLike(like)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "取消点赞失败",
		}).Error("取消点赞失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "取消点赞失败",
		})
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "取消点赞成功",
	})
}
