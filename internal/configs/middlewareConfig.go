package configs

import (
	"BBBingyan/internal/models/infoModels"
	"BBBingyan/internal/utils"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitMiddleware(e *echo.Echo) {

	//recover
	e.Use(middleware.Recover())

	//logger
	logFileLocation := filepath.Join("./logs", "httpLog.log")
	logFile, err := os.OpenFile(logFileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		utils.Log.WithFields(logrus.Fields{
			"error":         err,
			"error_message": "打开httpLog文件失败",
		}).Panic("打开httpLog文件失败")
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: "time:${time_rfc3339}\nmethod:${method}\nuri:${uri}\nstatus:${status}\nerror:${error}\nhost:${host}\npath:${path}\n\n\n",
	}))

	//csrf
	//配置文件区分开发和生产环境
	cookieSecure := viper.GetBool("CSRF.cookieSecure")
	cookieHTTPOnly := viper.GetBool("CSRF.cookieHTTPOnly")
	cookieMaxAge := viper.GetInt("CSRF.cookieMaxAge")
	csrfConfig := middleware.CSRFConfig{
		TokenLookup:    "header:X-CSRF-Token",
		CookiePath:     "/",
		CookieSecure:   cookieSecure,
		CookieHTTPOnly: cookieHTTPOnly,
		ContextKey:     "csrf",
		CookieMaxAge:   cookieMaxAge,
	}
	e.Use(middleware.CSRFWithConfig(csrfConfig))

	//JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			if (c.Path() == "/csrf-token" && c.Request().Method == "GET") || (c.Path() == "/users/account/activation/:activationCode" && c.Request().Method == "GET") || (c.Path() == "/users/account" && c.Request().Method == "POST") || (c.Path() == "/users/login" && c.Request().Method == "POST") {
				return true
			}
			return false
		},
		SigningKey:  []byte(viper.GetString("jwt.jwtSecret")),
		TokenLookup: "header:Authorization:Bearer ",
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*infoModels.JwtCustomClaim)
			utils.Log.WithFields(logrus.Fields{
				"userId":          claims.UserId,
				"isAdmin":         claims.IsAdmin,
				"success_message": "用户登录成功",
			}).Info("用户登录成功")
			c.Set("userId", claims.UserId)
			c.Set("isAdmin", claims.IsAdmin)
		},
	}))
}
