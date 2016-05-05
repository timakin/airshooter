package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")

	notifications := api.Group("/notifications")
	notifications.POST("/enqueue", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST enqueue")
	})
	notifications.PUT("/publish", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT publish")
	})
	notifications.GET("/subscriptions", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET subscriptions")
	})
	notifications.PUT("/subscribe", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT subscribe")
	})

	messages := api.Group("/messages")
	messages.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST messages")
	})
	messages.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET messages")
	})
	messages.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET message "+c.Param("id"))
	})

	e.Run(standard.New(":3000"))
}
