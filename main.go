package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs := config.GetConfig()
	db := _utils.InitDB(configs)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())
}
