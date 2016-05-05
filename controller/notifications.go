package controller

import (
	"github.com/labstack/echo"
	m "github.com/timakin/airshooter/model"
	"net/http"
)

func EnqueueNotification(c echo.Context) error {
	n := new(m.Notification)
	title := "test title"
	n.Title = &title
	return c.JSON(http.StatusOK, n)
}

func PublishNotification(c echo.Context) error {
	return c.String(http.StatusOK, "PUT publish")
}

func GetNotifications(c echo.Context) error {
	return c.String(http.StatusOK, "GET subscriptions")
}

func MarkReadNotifications(c echo.Context) error {
	return c.String(http.StatusOK, "PUT subscribe")
}
