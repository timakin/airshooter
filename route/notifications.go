package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
	"github.com/labstack/echo/middleware"
)

func AddNotificationAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/notifications")
	combined.Use(middleware.JWT([]byte("secret")))
	combined.Use(controller.ValidateToken)
	combined.POST("/enqueue", controller.EnqueueNotification)
	combined.GET("/:id", controller.GetNotification)
	combined.PUT("/read", controller.MarkAllNotificationsAsRead)
	combined.GET("/all", controller.GetAllNotifications)
	combined.GET("/unread", controller.GetUnreadNotifications)
	return combined
}
