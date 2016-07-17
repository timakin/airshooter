package route

import (
	"github.com/echo-contrib/echopprof"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/timakin/airshooter/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	AddClientAPI(api)
	AddNotificationAPI(api)
	AddMessageAPI(api)

	echopprof.Wrapper(e)

	return e
}
