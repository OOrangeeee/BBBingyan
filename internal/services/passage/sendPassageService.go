package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/dataModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"
	"strings"

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
	timeTool := utils.TimeTool{}
	passageTime, err := timeTool.GetCurrentTime()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取当前时间失败",
		}).Error("获取当前时间失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取当前时间失败",
		})
	}
	passage := &dataModels.Passage{
		PassageTitle:          passageTatle,
		PassageContent:        passageContent,
		PassageAuthorUserName: userNow.UserName,
		PassageAuthorNickName: userNow.UserNickName,
		PassageAuthorId:       userNow.ID,
		PassageTag:            passageTag,
		PassageTime:           passageTime,
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
	matchTool := utils.MatchTool{}
	atUsers := matchTool.MatchAtInString(passageContent)
	for _, atUser := range atUsers {
		if userMapper.IfUserExist(atUser) {
			userBeAt, err := userMapper.GetUsersByUserName(atUser)
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "获取被@用户失败",
				}).Error("获取被@用户失败")
				continue
			}
			if len(userBeAt) == 0 {
				utils.Log.WithField("error_message", "被@用户不存在").Error("被@用户不存在")
				continue
			}
			userBeAtNow := userBeAt[0]
			noticeEmailSubject := viper.GetString("email.emailOfAt.subject")
			noticeEmailContent := viper.GetString("email.emailOfAt.body")
			noticeEmailContent = strings.Replace(noticeEmailContent, "{用户名}", userBeAtNow.UserNickName, -1)
			noticeEmailContent = strings.Replace(noticeEmailContent, "{提到内容}", passageTatle, -1)
			noticeEmailContent = strings.Replace(noticeEmailContent, "{passage-id}", strconv.Itoa(int(passage.ID)), -1)
			noticeEmailContent = strings.Replace(noticeEmailContent, "{联系电话}", viper.GetString("info.contactPhone"), -1)
			noticeEmailContent = strings.Replace(noticeEmailContent, "{电子邮件地址}", viper.GetString("info.emailAddress"), -1)
			noticeEmailContent = strings.Replace(noticeEmailContent, "{官方网站}", viper.GetString("info.website"), -1)
			mileTool := utils.MileTool{}
			err = mileTool.SendMail([]string{userBeAtNow.UserEmail}, noticeEmailSubject, noticeEmailContent, viper.GetString("email.emailFromNickname"))
			if err != nil {
				utils.Log.WithFields(logrus.Fields{
					"error":         err,
					"error_message": "发送邮件失败",
				}).Error("发送邮件失败")
			}
		}
	}
	userNow.UserPassageCount++
	err = userMapper.UpdateUser(userNow)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "更新用户文章数失败",
		}).Error("更新用户文章数失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "更新用户文章数失败",
		})
	}
	followMapper := mappers.FollowMapper{}
	follows, err := followMapper.GetFollowsByToUserId(userId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取关注失败",
		}).Error("获取关注失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取关注失败",
		})
	}
	for _, follow := range follows {
		user, err := userMapper.GetUsersByUserId(follow.FromUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "获取关注用户失败",
			}).Error("获取关注用户失败")
			continue
		}
		if len(user) == 0 {
			utils.Log.WithField("error_message", "关注用户不存在").Error("关注用户不存在")
			continue
		}
		userNow := user[0]
		noticeEmailSubject := viper.GetString("email.emailOfFollow.subject")
		noticeEmailContent := viper.GetString("email.emailOfFollow.body")
		noticeEmailContent = strings.Replace(noticeEmailContent, "{用户名}", userNow.UserNickName, -1)
		noticeEmailContent = strings.Replace(noticeEmailContent, "{文章标题}", passageTatle, -1)
		noticeEmailContent = strings.Replace(noticeEmailContent, "{passage-id}", strconv.Itoa(int(passage.ID)), -1)
		noticeEmailContent = strings.Replace(noticeEmailContent, "{联系电话}", viper.GetString("info.contactPhone"), -1)
		noticeEmailContent = strings.Replace(noticeEmailContent, "{电子邮件地址}", viper.GetString("info.emailAddress"), -1)
		noticeEmailContent = strings.Replace(noticeEmailContent, "{官方网站}", viper.GetString("info.website"), -1)
		mileTool := utils.MileTool{}
		err = mileTool.SendMail([]string{userNow.UserEmail}, noticeEmailSubject, noticeEmailContent, viper.GetString("email.emailFromNickname"))
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "发送邮件失败",
			}).Error("发送邮件失败")
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "发布文章成功",
	})
}
