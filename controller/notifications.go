package controller

import (
	"github.com/labstack/echo"
	m "github.com/timakin/airshooter/model"
	//	s "github.com/timakin/airshooter/service"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func EnqueueNotification(c echo.Context) error {
	notification := new(m.Notification)
	if err := c.Bind(notification); err != nil {
		return err
	}

	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if err := validate.Struct(notification); err != nil {
		fmt.Println(err)
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
