package route

import (	
	"os"
	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
	"github.com/labstack/echo/middleware"
)

func AddMessageAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/messages")
	combined.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SIGN_KEY"))))
	combined.Use(controller.ValidateToken)
	combined.POST("/send", controller.SendMessage)
	combined.GET("/received", controller.ReceiveMessages)
	combined.GET("/threads", controller.GetThreads)
	return combined
}
