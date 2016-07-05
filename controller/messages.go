package controller

import (
	"github.com/labstack/echo"
	"github.com/timakin/airshooter/constant"
	m "github.com/timakin/airshooter/model"
	s "github.com/timakin/airshooter/service"
	"net/http"
	"strconv"
)

func SendMessage(c echo.Context) error {
	message := new(m.Message)
	if err := c.Bind(message); err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrInternalServerError)
	}

	if err := ValidationHandler(message); err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrRequestInvalid)
	}

	result, err := s.InsertMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func ReceiveMessages(c echo.Context) error {
	senderId, _ := strconv.ParseInt(c.QueryParam("senderId"), 10, 64)
	recipientId, _ := strconv.ParseInt(c.QueryParam("recipientId"), 10, 64)
	params := map[string]interface{}{"senderId": &senderId, "recipientId": &recipientId}
	result, err := s.GetMessages(&params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func ListThreads(c echo.Context) error {
	recipientId, _ := strconv.ParseInt(c.QueryParam("recipientId"), 10, 64)
	params := map[string]interface{}{"recipientId": &recipientId}
	result, err := s.GetMessages(&params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}
