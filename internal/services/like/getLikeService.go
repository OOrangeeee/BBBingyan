package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetUserLikesByUserIdService(paramsMap map[string]string, c echo.Context) error {
	likeMapper := mappers.LikeMapper{}
	userMapper := mappers.UserMapper{}
	passageMapper := mappers.PassageMapper{}
	fromUserIdStr := paramsMap["fromUserId"]
	if fromUserIdStr == "" {
		utils.Log.WithField("error_message", "用户ID为空").Error("用户ID为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户ID为空",
		})
	}
	fromUserIdUint64, err := strconv.ParseUint(fromUserIdStr, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "用户ID转换失败",
		}).Error("用户ID转换失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户ID转换失败",
		})
	}
	fromUserId := uint(fromUserIdUint64)
	if !userMapper.IfUserExistById(fromUserId) {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	users, err := userMapper.GetUsersByUserId(fromUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		}).Error("获取用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	if len(users) == 0 {
		utils.Log.WithField("error_message", "用户不存在").Error("用户不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	userInfo := infoModels.User{
		UserId:           users[0].ID,
		UserName:         users[0].UserName,
		UserEmail:        users[0].UserEmail,
		UserNickName:     users[0].UserNickName,
		UserFollowCount:  users[0].UserFollowCount,
		UserFansCount:    users[0].UserFansCount,
		UserPassageCount: users[0].UserPassageCount,
		UserLikeCount:    users[0].UserLikeCount,
		UserIsAdmin:      users[0].UserIsAdmin,
	}
	likes, err := likeMapper.GetLikesByFromUserId(fromUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户点赞失败",
		}).Error("获取用户点赞失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户点赞失败",
		})
	}
	var ansLikes []infoModels.Like
	for _, like := range likes {
		passages, err := passageMapper.GetPassagesByID(like.ToPassageId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "获取文章失败",
			}).Error("获取文章失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "获取文章失败",
			})
		}
		if len(passages) == 0 {
			utils.Log.WithField("error_message", "文章不存在").Error("文章不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "文章不存在",
			})
		}
		passageInfo := infoModels.PassageBrief{
			ID:                    passages[0].ID,
			PassageTitle:          passages[0].PassageTitle,
			PassageAuthorUserName: passages[0].PassageAuthorUserName,
			PassageAuthorNickName: passages[0].PassageAuthorNickName,
			PassageTag:            passages[0].PassageTag,
			PassageBeLikedCount:   passages[0].PassageBeLikedCount,
			PassageCommentCount:   passages[0].PassageCommentCount,
			PassageTime:           passages[0].PassageTime,
		}
		likeInfo := infoModels.Like{
			ID:        like.ID,
			FromUser:  userInfo,
			ToPassage: passageInfo,
		}
		ansLikes = append(ansLikes, likeInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"likes":           ansLikes,
		"success_message": "获取用户点赞成功",
	})
}
