package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddClientAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/clients")
	g.Use(middleware.BasicAuth(func(clientId, clientSecret string) bool {
		if clientId == "joe" && clientSecret == "secret" {
			return true
		}
		return false
	}))
	combined.POST("/authenticate", controller.Authenticate)
	return combined
}
