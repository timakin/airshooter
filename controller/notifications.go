package controller

import (
    "github.com/timakin/airshooter/constant"
	"github.com/labstack/echo"
    "github.com/pkg/errors"
	"net/http"
)

func EnqueueNotification(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // service
        // if err.. return err
        return next(c)
    }
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
