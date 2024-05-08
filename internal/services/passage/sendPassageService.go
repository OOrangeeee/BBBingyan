package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SendPassageService(paramsMap map[string]string, c echo.Context) error {
	userMapper := mappers.UserMapper{}
	passageTatle := paramsMap["passageTitle"]
	passageContent := paramsMap["passageContent"]
	passageTag := paramsMap["passageTag"]
	if passageTatle == "" || passageContent == "" || passageTag == "" {
		utils.Log.WithField("error_message", "文章信息不完整").Error("文章信息不完整")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章信息不完整",
		})
	}
	tagString := viper.GetString("passage.tags")
	tags := strings.Split(tagString, ",")
	tagExist := false
	for _, tag := range tags {
		if tag == passageTag {
			tagExist = true
			break
		}
	}
	if !tagExist {
		utils.Log.WithField("error_message", "文章标签不存在").Error("文章标签不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章标签不存在",
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
	passage := &dataModels.Passage{
		PassageTitle:          passageTatle,
		PassageContent:        passageContent,
		PassageAuthorUserName: userNow.UserName,
		PassageAuthorNickName: userNow.UserNickName,
		PassageAuthorId:       userNow.ID,
		PassageTag:            passageTag,
		PassageTime:           time.Now(),
	}
	passageMapper := mappers.PassageMapper{}
	err = passageMapper.AddNewPassage(passage)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "发布文章失败",
		}).Error("发布文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "发布文章失败",
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
		"success_message": "发布文章成功",
	})
}
