package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/timakin/airshooter/controller"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")

	notifications := api.Group("/notifications")
	notifications.POST("/enqueue", controller.EnqueueNotification)
	notifications.PUT("/publish", controller.PublishNotification)
	notifications.GET("/subscriptions", controller.GetNotifications)
	notifications.PUT("/subscriptions", controller.MarkReadNotifications)

	messages := api.Group("/messages")
	messages.POST("", controller.PostMessage)
	messages.GET("", controller.GetMessages)

	e.Run(standard.New(":3000"))
}
