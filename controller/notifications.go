package controller

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
	s "github.com/timakin/airshooter/service"
	"net/http"
)

func EnqueueNotification(c echo.Context) error {
	notification := new(NotificationPostRequest)
	if err := c.Bind(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(notification); err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	if result, err := s.EnqueueNotification(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func PublishNotification(c echo.Context) error {
	return c.String(http.StatusOK, "PUT publish")
}

func GetNotification(c echo.Context) error {
	notification := new(NotificationGetRequest)
	if err := c.Bind(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(notification); err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	if result, err := s.GetNotification(notification.Id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetNotifications(c echo.Context) error {
	if result, err := s.GetNotifications(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func MarkReadNotifications(c echo.Context) error {
	return c.String(http.StatusOK, "PUT subscribe")
}
