package route

import (
	"github.com/echo-contrib/echopprof"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/timakin/airshooter/controller"
)

func Init() *echo.Echo {
	e := echo.New()
	DefaultJWTConfig := JWTConfig{
		SigningMethod: AlgorithmHS256,
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
	}
	e.Use(middleware.JWTWithConfig(DefaultJWTConfig))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(controller.AuthClient)
	e.Use(controller.ValidateRequest)

	api := e.Group("/api")
	AddNotificationAPI(api)
	AddMessageAPI(api)

	echopprof.Wrapper(e)

	return e
}
