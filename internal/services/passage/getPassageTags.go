package services

import (
	"BBBingyan/internal/models/infoModels"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func GetPassageTags(c echo.Context) error {
	tagString := viper.GetString("passage.tags")
	tags := strings.Split(tagString, ",")
	tag := infoModels.Tags(tags)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"tags":            tag,
		"success_message": "获取文章标签成功",
	})
}
