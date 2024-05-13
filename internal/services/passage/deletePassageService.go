package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func DeletePassageService(paramaMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageId := paramaMap["passageId"]
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "文章ID转换失败",
		})
	}
	passageIdUint := uint(passageIdUint64)
	toPassage, err := passageMapper.GetPassagesByID(passageIdUint)
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
	userId := c.Get("userId").(uint)
	if toPassage[0].PassageAuthorId != userId {
		utils.Log.WithField("error_message", "无权删除文章").Error("无权删除文章")
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"error_message": "无权删除文章",
		})
	}
	err = passageMapper.DeletePassage(toPassage[0])
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "删除文章失败",
		}).Error("删除文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "删除文章失败",
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
		"success_message": "删除文章成功",
	})
}
