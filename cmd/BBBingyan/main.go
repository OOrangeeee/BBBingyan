package main

import (
	"BBBingyan/internal/configs"
	"BBBingyan/internal/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	utils.InitLog()
	configs.InitViper()
	utils.InitDB()
	configs.InitMiddleware(e)
	configs.GetRouterConfig(e)
	configs.PostRouterConfig(e)
	configs.PutRouterConfig(e)
	configs.DeleteRouterConfig(e)
	e.Logger.Fatal(e.Start(":714"))
}
