package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddMessageAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/messages")
	combined.POST("", controller.PostMessage)
	combined.GET("", controller.GetMessages)
	return combined
}
