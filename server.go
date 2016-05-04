package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/api/notifications/enqueue", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST enqueue")
	})
	e.PUT("/api/notifications/publish", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT publish")
	})
	e.GET("/api/notifications/subscriptions", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET subscriptions")
	})
	e.PUT("/api/notifications/subscribe", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT subscribe")
	})

	e.POST("/api/messages", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST messages")
	})
	e.GET("/api/messages", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET messages")
	})
	e.GET("/api/messages/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET message"+c.Param("id"))
	})

	e.Run(standard.New(":3000"))
}
