package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddMessageAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/messages")
	combined.POST("/send", controller.SendMessage)
	combined.GET("/receive", controller.ReceiveMessages)
	return combined
}
