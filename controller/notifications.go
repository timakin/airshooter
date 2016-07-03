package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
	s "github.com/timakin/airshooter/service"
	"net/http"
	"strconv"
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
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetNotification(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	result, err := s.GetNotification(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllNotifications(c echo.Context) error {
	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 64)
	params := map[string]interface{}{"recipientId": &userId}
	result, err := s.GetNotifications(&params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func GetUnreadNotifications(c echo.Context) error {
	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 64)
	params := map[string]interface{}{"recipientId": &userId, "status": "unread"}
	result, err := s.GetNotifications(&params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func MarkAllNotificationsAsRead(c echo.Context) error {
	userId, _ := strconv.ParseInt(c.QueryParam("userId"), 10, 64)
	err := s.MarkNotificationsAsRead(&userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
