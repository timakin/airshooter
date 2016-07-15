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

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	//	e.Use(controller.AuthClient)
	e.Use(controller.ValidateRequest)

	api := e.Group("/api")
	AddNotificationAPI(api)
	AddMessageAPI(api)

	echopprof.Wrapper(e)

	return e
}
