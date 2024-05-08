package services

import (
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetPassageTagsService(c echo.Context) error {
	tagString := viper.GetString("passage.tags")
	tags := strings.Split(tagString, ",")
	tag := infoModels.Tags(tags)
	csrfTool := utils.CSRFTool{}
	getCSRF := csrfTool.SetCSRFToken(c)
	if !getCSRF {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "CSRF Token 获取失败",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"tags":            tag,
		"success_message": "获取文章标签成功",
	})
}
