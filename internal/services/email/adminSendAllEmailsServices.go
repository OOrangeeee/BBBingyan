package services

import (
	"BBBingyan/internal/mappers"
	"BBBingyan/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SendAllEmails(paramsMap map[string]string, c echo.Context) error {
	ifAdmin := c.Get("isAdmin").(bool)
	if !ifAdmin {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error_message": "没有权限",
		})
	}
	subject := paramsMap["subject"]
	body := paramsMap["body"]
	fromNickName := paramsMap["fromNickName"]
	if subject == "" || body == "" || fromNickName == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error_message": "邮件主题、内容或发件人昵称为空",
		})
	}
	userEmailMapper := mappers.UserEmailMapper{}
	userEmails, err := userEmailMapper.GetAllUserEmails()
	if err != nil {
		utils.Log.WithField("error_message", "获取所有用户邮箱失败").Error("获取所有用户邮箱失败")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error_message": "获取所有用户邮箱失败",
		})
	}
	for _, email := range userEmails {
		emailAddress := email.Email
		go func(emailAddress string) {
			utils.Sender.EmailQueue <- utils.EmailTask{
				Recipient:    emailAddress,
				Subject:      subject,
				Body:         body,
				FromNickName: fromNickName,
			}
		}(emailAddress)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success_message": "所有邮件发送完成",
	})
}
