package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/timakin/airshooter/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(controller.AuthClient)
	e.Use(controller.ValidateRequest)

	api := e.Group("/api")
	AddNotificationAPI(api)
	AddMessageAPI(api)
	return e
}
