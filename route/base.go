package route

import (
	"github.com/labstack/echo/middleware"
    "github.com/labstack/echo"
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

