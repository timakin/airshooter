package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type Notification struct {
}

type Notifications []Notification

func EnqueueNotification(c echo.Context) error {
	return c.String(http.StatusOK, "POST enqueue")
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
