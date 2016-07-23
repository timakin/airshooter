package route

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
	"github.com/labstack/echo/middleware"
)

func AddNotificationAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/notifications")
	combined.Use(middleware.JWT([]byte(os.Getenv("TOKEN_SIGN_KEY"))))
	combined.Use(controller.ValidateToken)
	combined.POST("/enqueue", controller.EnqueueNotification)
	combined.GET("/:id", controller.GetNotification)
	combined.PUT("/read", controller.MarkAllNotificationsAsRead)
	combined.GET("/all", controller.GetAllNotifications)
	combined.GET("/unread", controller.GetUnreadNotifications)
	return combined
}
