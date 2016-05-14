package controller

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
	//	s "github.com/timakin/airshooter/service"
	"net/http"
)

func EnqueueNotification(c echo.Context) error {
	notification := new(m.Notification)
	if err := c.Bind(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(notification); err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	return c.JSON(http.StatusCreated, notification)
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
