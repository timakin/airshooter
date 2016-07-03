package controller

import (
	"github.com/labstack/echo"
	m "github.com/timakin/airshooter/model"
	"net/http"
)

func SendMessage(c echo.Context) error {
	return c.String(http.StatusOK, "POST messages")
}

func ReceiveMessages(c echo.Context) error {
	return c.JSON(http.StatusOK, message)
}
