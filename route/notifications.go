package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddNotificationAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/notifications")
	combined.POST("/enqueue", controller.EnqueueNotification)
	combined.GET("/:id", controller.GetNotification)
	combined.PUT("/read", controller.MarkAllNotificationsAsRead)
	combined.GET("/all", controller.GetAllNotifications)
	combined.PUT("/unread", controller.GetUnreadNotifications)
	return combined
}
