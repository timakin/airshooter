package route

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/controller"
)

func AddNotificationAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/notifications")
	combined.POST("/enqueue", controller.EnqueueNotification)
	combined.GET("/subscriptions/:id", controller.GetNotification)
	// combined.PUT("/publish", controller.PublishNotification)
	// combined.GET("/subscriptions", controller.GetNotifications)
	// combined.PUT("/subscriptions", controller.MarkReadNotifications)
	return combined
}
