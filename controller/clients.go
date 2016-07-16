package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name":  "Jon Snow",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
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
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
