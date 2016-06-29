package controller

import (
	"fmt"
	"github.com/k0kubun/pp"
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
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetNotification(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	pp.Print(id)
	result, err := s.GetNotification(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func PublishNotifications(c echo.Context) error {
	return c.String(http.StatusOK, "PUT publish")
}

func GetNotifications(c echo.Context) error {
	result, err := s.GetNotifications()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

// func MarkReadNotifications(c echo.Context) error {
// 	return c.String(http.StatusOK, "PUT subscribe")
// }
