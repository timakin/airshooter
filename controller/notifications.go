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
