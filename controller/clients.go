package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Authenticate(c echo.Context) error {
	// Basic認証スキームを通してclient_id, client_secretをAPIに渡す。
	// DBのclientsテーブルに保存されたclient_id & passwordと照合して、
	// 正規のクライアントであればtokenを引き渡す
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"client_id": clientId,
			"admin":     true,
			"exp":       time.Now().Add(time.Hour * 72).Unix(),
		})
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	clientId := claims["client_id"].(string)
	return c.String(http.StatusOK, "Welcome "+clientId+"!")
}
