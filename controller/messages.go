package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type Message struct {
	Text      string `json: "text"`
	From      string
	To        string
	ThreadId  int
	CreatedAt int
}

func PostMessage(c echo.Context) error {
	return c.String(http.StatusOK, "POST messages")
}

func GetMessages(c echo.Context) error {
	m := new(Message)
	m.Text = "test"
	return c.JSON(http.StatusOK, m)
}
