package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	AddNotificationAPI(api)
	AddMessageAPI(api)
	return e
}
