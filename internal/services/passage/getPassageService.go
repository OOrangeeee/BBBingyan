package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetPassageByIDService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	id := paramsMap["id"]
	if id == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "id为空",
		}).Error("id为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "id为空",
		})
	}
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析id失败",
		}).Error("解析id失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析id失败",
		})
	}
	idUint := uint(idUint64)
	passages, err := passageMapper.GetPassagesByID(idUint)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	passage := passages[0]
	passageInfo := infoModels.PassageDetail{
		ID:                    passage.ID,
		PassageTitle:          passage.PassageTitle,
		PassageContent:        passage.PassageContent,
		PassageAuthorUserName: passage.PassageAuthorUserName,
		PassageAuthorNickName: passage.PassageAuthorNickName,
		PassageAuthorId:       passage.PassageAuthorId,
		PassageTag:            passage.PassageTag,
		PassageBeLikedCount:   passage.PassageBeLikedCount,
		PassageCommentCount:   passage.PassageCommentCount,
		PassageTime:           passage.PassageTime,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfo":     passageInfo,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageTitleService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageTitle := paramsMap["passageTitle"]
	if passageTitle == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章标题为空",
		}).Error("文章标题为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章标题为空",
		})
	}
	pageSize := paramsMap["pageSize"]
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	passages, err := passageMapper.GetPassagesByPassageTitle(passageTitle)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageAuthorUserNameService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageAuthorUserName := paramsMap["passageAuthorUserName"]
	if passageAuthorUserName == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "作者用户名为空",
		}).Error("作者用户名为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "作者用户名为空",
		})
	}
	pageSize := paramsMap["pageSize"]
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	passages, err := passageMapper.GetPassagesByPassageAuthorUserName(passageAuthorUserName)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageAuthorNickNameService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageAuthorNickName := paramsMap["passageAuthorNickName"]
	if passageAuthorNickName == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "作者昵称为空",
		}).Error("作者昵称为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "作者昵称为空",
		})
	}
	pageSize := paramsMap["pageSize"]
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	passages, err := passageMapper.GetPassagesByPassageAuthorNickName(passageAuthorNickName)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageAuthorIdService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageAuthorIdUint := c.Get("userId").(uint)
	pageSize := paramsMap["pageSize"]
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	passages, err := passageMapper.GetPassagesByPassageAuthorId(passageAuthorIdUint)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageTagService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageTag := paramsMap["passageTag"]
	if passageTag == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章标签为空",
		}).Error("文章标签为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章标签为空",
		})
	}
	pageSize := paramsMap["pageSize"]
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	tagString := viper.GetString("passage.tags")
	tags := strings.Split(tagString, ",")
	isTag := false
	for _, tag := range tags {
		if tag == passageTag {
			isTag = true
			break
		}
	}
	if !isTag {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章标签不存在",
		}).Error("文章标签不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章标签不存在",
		})
	}
	passages, err := passageMapper.GetPassagesByPassageTag(passageTag)
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
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

// GetHottestPassagesService 获取最火的文章
func GetHottestPassagesService(c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passages, err := passageMapper.GetAllPassages()
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取文章失败",
		}).Error("获取文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取文章失败",
		})
	}
	// 计算每篇文章的得分
	var scores []struct {
		ID    uint
		Score float64
	}
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	maxTime := time.Now().Sub(epoch).Seconds() // 获取最大时间差，用于归一化
	for _, passage := range passages {
		// 文章发布时间的秒数
		duration := passage.PassageTime.Sub(epoch).Seconds()
		// 归一化
		normalizedTime := duration / maxTime
		// 使用指数衰减公式来计算时间权重。公式中的-normalizedTime * 10表示随着时间的推移，权重以指数级别减小。
		timeWeight := math.Exp(-normalizedTime * 10) // 使用指数衰减函数
		// 结合文章的点赞数和评论数（评论数乘以2，表示评论通常比点赞更能表明用户的参与度和对内容的兴趣），并乘以时间权重。
		scoreT := float64(passage.PassageBeLikedCount) + float64(passage.PassageCommentCount*2)*timeWeight
		scores = append(scores, struct {
			ID    uint
			Score float64
		}{passage.ID, scoreT})
	}

	// 根据得分排序
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// 准备返回前十篇文章的ID
	var topArticleIDs []uint
	topArticleIDs = make([]uint, 11)
	for _, score := range scores {
		topArticleIDs = append(topArticleIDs, score.ID)
		if len(topArticleIDs) >= 10 { // 如果只需要前10篇文章
			break
		}
	}

	var resPassage []infoModels.PassageBrief
	for _, passageID := range topArticleIDs {
		passageTemp, err := passageMapper.GetPassagesByID(passageID)
		if err != nil {
			utils.Log.WithFields(logrus.Fields{
				"error":         err,
				"error_message": "获取文章失败",
			}).Error("获取文章失败")
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error_message": "获取文章失败",
			})
		}
		if len(passageTemp) == 0 {
			utils.Log.WithFields(logrus.Fields{
				"error_message": "文章不存在",
			}).Error("文章不存在")
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error_message": "文章不存在",
			})
		}
		passage := passageTemp[0]
		passageBrief := infoModels.PassageBrief{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageBeLikedCount:   passage.PassageBeLikedCount,
			PassageCommentCount:   passage.PassageCommentCount,
			PassageTime:           passage.PassageTime,
		}
		resPassage = append(resPassage, passageBrief)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    resPassage,
		"success_message": "获取文章成功",
	})
}

func SearchPassagesService(paramsMap map[string]string, c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	searchTitle := paramsMap["searchTitle"]
	// 一页显示的文章数量
	pageSize := paramsMap["pageSize"]
	if searchTitle == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "搜索标题为空",
		}).Error("搜索标题为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "搜索标题为空",
		})
	}
	if pageSize == "" {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小为空",
		}).Error("页大小为空")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小为空",
		})
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "解析页大小失败",
		}).Error("解析页大小失败")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "解析页大小失败",
		})
	}
	if pageSizeInt <= 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "页大小小于等于0",
		}).Error("页大小小于等于0")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "页大小小于等于0",
		})
	}
	passages, err := passageMapper.SearchPassagesByPassageTitle(searchTitle)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "搜索文章失败",
		}).Error("搜索文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "搜索文章失败",
		})
	}
	if len(passages) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章不存在",
		}).Error("文章不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos [][]infoModels.PassageBrief
	// 将文章信息存到passageInfo按页分组
	for i := 0; i < len(passages); i += pageSizeInt {
		var passageInfo []infoModels.PassageBrief
		for j := i; j < i+pageSizeInt && j < len(passages); j++ {
			passage := passages[j]
			passageBrief := infoModels.PassageBrief{
				ID:                    passage.ID,
				PassageTitle:          passage.PassageTitle,
				PassageAuthorUserName: passage.PassageAuthorUserName,
				PassageAuthorNickName: passage.PassageAuthorNickName,
				PassageTag:            passage.PassageTag,
				PassageBeLikedCount:   passage.PassageBeLikedCount,
				PassageCommentCount:   passage.PassageCommentCount,
				PassageTime:           passage.PassageTime,
			}
			passageInfo = append(passageInfo, passageBrief)
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "搜索文章成功",
	})
}
