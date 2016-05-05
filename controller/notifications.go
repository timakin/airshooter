package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

// from, to = id, type
// CreatedAt, Expiry = unixtimestamp
type Notification struct {
	Title     string `json: "title"`
	Text      string `json: "text"`
	From      string `json: "from"`
	To        string `json: "to"`
	CreatedAt int    `json: "created_at"`
	Expiry    int    `json: "expiry"`
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
