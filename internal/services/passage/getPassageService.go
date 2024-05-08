package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"net/http"
	"sort"
	"strconv"
	"strings"

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
	passage, err := passageMapper.GetPassageByID(idUint)
	passageInfo := infoModels.Passage{
		ID:                    passage.ID,
		PassageTitle:          passage.PassageTitle,
		PassageContent:        passage.PassageContent,
		PassageAuthorUserName: passage.PassageAuthorUserName,
		PassageAuthorNickName: passage.PassageAuthorNickName,
		PassageTag:            passage.PassageTag,
		PassageTime:           passage.PassageTime,
	}
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "获取文章失败",
		}).Error("获取文章失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取文章失败",
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
	var passageInfos []infoModels.Passage
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
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
	var passageInfos []infoModels.Passage
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
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
	var passageInfos []infoModels.Passage
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

func GetPassagesByPassageAuthorIdService(c echo.Context) error {
	passageMapper := mappers.PassageMapper{}
	passageAuthorIdUint := c.Get("userId").(uint)
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
	var passageInfos []infoModels.Passage
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
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
	tagString := viper.GetString("passage.tag")
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
	var passageInfos []infoModels.Passage
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}

// 获取最近的五个文章
func GetLastPassagesService(c echo.Context) error {
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
	if len(passages) == 0 {
		utils.Log.WithFields(logrus.Fields{
			"error_message": "文章不存在",
		}).Error("文章不存在")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "文章不存在",
		})
	}
	var passageInfos []infoModels.Passage
	sort.Slice(passages, func(i, j int) bool {
		return passages[i].PassageTime.After(passages[j].PassageTime)
	})
	if len(passages) > 5 {
		passages = passages[:5]
	}
	for _, passage := range passages {
		passageInfo := infoModels.Passage{
			ID:                    passage.ID,
			PassageTitle:          passage.PassageTitle,
			PassageContent:        passage.PassageContent,
			PassageAuthorUserName: passage.PassageAuthorUserName,
			PassageAuthorNickName: passage.PassageAuthorNickName,
			PassageTag:            passage.PassageTag,
			PassageTime:           passage.PassageTime,
		}
		passageInfos = append(passageInfos, passageInfo)
	}
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"passageInfos":    passageInfos,
		"success_message": "获取文章成功",
	})
}
