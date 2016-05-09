package route

import (
    "github.com/timakin/airshooter/controller"
    "github.com/labstack/echo"
)

func AddMessageAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/messages")
    combined.POST("", controller.PostMessage)
	combined.GET("", controller.GetMessages)
    return combined
}

