package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
	s "github.com/timakin/airshooter/service"
	"net/http"
)

func EnqueueNotification(c echo.Context) error {
	notification := new(m.Notification)
	if err := c.Bind(notification); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(notification); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	result, err := s.EnqueueNotification(notification)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func PublishNotification(c echo.Context) error {
	return c.String(http.StatusOK, "PUT publish")
}

func GetNotification(c echo.Context) error {
	var notification *m.Notification
	if err := c.Bind(notification); err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(notification); err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	result, err := s.GetNotification(notification.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetNotifications(c echo.Context) error {
	result, err := s.GetNotifications()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func MarkReadNotifications(c echo.Context) error {
	return c.String(http.StatusOK, "PUT subscribe")
}
