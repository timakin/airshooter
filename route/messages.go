package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
	"github.com/labstack/echo/middleware"
)

func AddMessageAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/messages")
	combined.Use(middleware.JWT([]byte("secret")))
	combined.Use(controller.ValidateToken)
	combined.POST("/send", controller.SendMessage)
	combined.GET("/received", controller.ReceiveMessages)
	combined.GET("/threads", controller.GetThreads)
	return combined
}
