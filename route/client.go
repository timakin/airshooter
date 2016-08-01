package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddClientAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/clients")
	combined.POST("/register", controller.Register, controller.BasicAuthForRegistration())
	combined.POST("/authenticate", controller.Authenticate)
	return combined
}
