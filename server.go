package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/timakin/airshooter/controller"
	"net/http"
)

type Message struct {
	Text string `json: "text"`
}

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
	messages.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST messages")
	})
	messages.GET("", func(c echo.Context) error {
		m := new(Message)
		m.Text = "test"
		return c.JSON(http.StatusOK, m)
	})
	messages.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET message "+c.Param("id"))
	})

	e.Run(standard.New(":3000"))
}
