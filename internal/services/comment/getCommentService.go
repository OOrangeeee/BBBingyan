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

func GetCommentByIdService(paramsMap map[string]string, c echo.Context) error {
	commentMapper := mappers.CommentMapper{}
	userMapper := mappers.UserMapper{}
	passageMapper := mappers.PassageMapper{}
	commentId := paramsMap["commentId"]
	if commentId == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "commentId为空",
		}).Error("commentId为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "commentId为空",
		})
	}
	commentIdUint64, err := strconv.ParseUint(commentId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析commentId失败",
		}).Error("解析commentId失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "解析commentId失败",
		})
	}
	commentUint := uint(commentIdUint64)
	comments, err := commentMapper.GetCommentById(commentUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取评论失败",
		}).Error("获取评论失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取评论失败",
		})
	}
	if len(comments) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "评论不存在",
		}).Error("评论不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "评论不存在",
		})
	}
	comment := comments[0]
	commentUserId := comment.FromUserId
	commentToPassageId := comment.ToPassageId
	commentContent := comment.CommentContent
	commentUsers, err := userMapper.GetUsersByUserId(commentUserId)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		}).Error("获取用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	if len(commentUsers) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "用户不存在",
		}).Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	commentUser := commentUsers[0]
	userInfo := infoModels.User{
		UserId:       commentUser.ID,
		UserName:     commentUser.UserName,
		UserEmail:    commentUser.UserEmail,
		UserNickName: commentUser.UserNickName,
		UserIsAdmin:  commentUser.UserIsAdmin,
	}
	passages, err := passageMapper.GetPassagesByID(commentToPassageId)
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
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章不存在",
		}).Error("文章不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	passage := passages[0]
	passageInfo := infoModels.PassageBrief{
		ID:                    passage.ID,
		PassageTitle:          passage.PassageTitle,
		PassageAuthorUserName: passage.PassageAuthorUserName,
		PassageAuthorNickName: passage.PassageAuthorNickName,
		PassageTag:            passage.PassageTag,
		PassageBeLikedCount:   passage.PassageBeLikedCount,
		PassageCommentCount:   passage.PassageCommentCount,
		PassageTime:           passage.PassageTime,
	}
	commentInfo := infoModels.Comment{
		CommentContent: commentContent,
		FromUser:       userInfo,
		ToPassage:      passageInfo,
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取评论成功",
		"commentInfo":     commentInfo,
	})
}

func GetCommentsByFromUserIdService(paramsMap map[string]string, c echo.Context) error {
	commentMapper := mappers.CommentMapper{}
	userMapper := mappers.UserMapper{}
	passageMapper := mappers.PassageMapper{}
	fromUserId := paramsMap["fromUserId"]
	if fromUserId == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "fromUserId为空",
		}).Error("fromUserId为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "fromUserId为空",
		})
	}
	fromUserIdUint64, err := strconv.ParseUint(fromUserId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析fromUserId失败",
		}).Error("解析fromUserId失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "解析fromUserId失败",
		})
	}
	fromUserIdUint := uint(fromUserIdUint64)
	comments, err := commentMapper.GetCommentsByFromUserId(fromUserIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取评论失败",
		}).Error("获取评论失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取评论失败",
		})
	}
	if len(comments) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "评论不存在",
		}).Error("评论不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "评论不存在",
		})
	}
	commentUsers, err := userMapper.GetUsersByUserId(fromUserIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取用户失败",
		}).Error("获取用户失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取用户失败",
		})
	}
	if len(commentUsers) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "用户不存在",
		}).Error("用户不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "用户不存在",
		})
	}
	commentUser := commentUsers[0]
	userInfo := infoModels.User{
		UserId:       commentUser.ID,
		UserName:     commentUser.UserName,
		UserEmail:    commentUser.UserEmail,
		UserNickName: commentUser.UserNickName,
		UserIsAdmin:  commentUser.UserIsAdmin,
	}
	var commentInfos []infoModels.Comment
	for _, comment := range comments {
		commentToPassageId := comment.ToPassageId
		commentContent := comment.CommentContent
		passages, err := passageMapper.GetPassagesByID(commentToPassageId)
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
			utils.Log.WithFields(logrus.Fields{
				"error_message": "文章不存在",
			}).Error("文章不存在")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error_message": "文章不存在",
			})
		}
		passage := passages[0]
		passageInfo := infoModels.PassageBrief{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageBeLikedCount:   passage.PassageBeLikedCount,
			PassageCommentCount:   passage.PassageCommentCount,
			PassageTime:           passage.PassageTime,
		}
		commentInfo := infoModels.Comment{
			CommentContent: commentContent,
			FromUser:       userInfo,
			ToPassage:      passageInfo,
		}
		commentInfos = append(commentInfos, commentInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取评论成功",
		"commentInfos":    commentInfos,
	})
}

func GetCommentsByToPassageIdService(paramsMap map[string]string, c echo.Context) error {
	commentMapper := mappers.CommentMapper{}
	userMapper := mappers.UserMapper{}
	passageMapper := mappers.PassageMapper{}
	toPassageId := paramsMap["toPassageId"]
	if toPassageId == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "toPassageId为空",
		}).Error("toPassageId为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "toPassageId为空",
		})
	}
	toPassageIdUint64, err := strconv.ParseUint(toPassageId, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析toPassageId失败",
		}).Error("解析toPassageId失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "解析toPassageId失败",
		})
	}
	toPassageIdUint := uint(toPassageIdUint64)
	comments, err := commentMapper.GetCommentsByToPassageId(toPassageIdUint)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取评论失败",
		}).Error("获取评论失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取评论失败",
		})
	}
	if len(comments) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "评论不存在",
		}).Error("评论不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "评论不存在",
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
	if len(passages) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章不存在",
		}).Error("文章不存在")
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	passage := passages[0]
	passageInfo := infoModels.PassageBrief{
		ID:                    passage.ID,
		PassageTitle:          passage.PassageTitle,
		PassageAuthorUserName: passage.PassageAuthorUserName,
		PassageAuthorNickName: passage.PassageAuthorNickName,
		PassageTag:            passage.PassageTag,
		PassageBeLikedCount:   passage.PassageBeLikedCount,
		PassageCommentCount:   passage.PassageCommentCount,
		PassageTime:           passage.PassageTime,
	}
	var commentInfos []infoModels.Comment
	for _, comment := range comments {
		commentUserId := comment.FromUserId
		commentContent := comment.CommentContent
		commentUsers, err := userMapper.GetUsersByUserId(commentUserId)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "获取用户失败",
			}).Error("获取用户失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "获取用户失败",
			})
		}
		if len(commentUsers) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "用户不存在",
			}).Error("用户不存在")
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error_message": "用户不存在",
			})
		}
		commentUser := commentUsers[0]
		userInfo := infoModels.User{
			UserId:       commentUser.ID,
			UserName:     commentUser.UserName,
			UserEmail:    commentUser.UserEmail,
			UserNickName: commentUser.UserNickName,
			UserIsAdmin:  commentUser.UserIsAdmin,
		}
		commentInfo := infoModels.Comment{
			CommentContent: commentContent,
			FromUser:       userInfo,
			ToPassage:      passageInfo,
		}
		commentInfos = append(commentInfos, commentInfo)

	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "获取评论成功",
		"commentInfos":    commentInfos,
	})
}
