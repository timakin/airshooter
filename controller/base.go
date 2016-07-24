package controller

import (
	s "github.com/timakin/airshooter/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v8"

	"log"
	"net/http"
	"time"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		// DBに登録されたtokenと照合して、validなclientかどうか検証する
		claims := user.Claims.(jwt.MapClaims)
		clientId := claims["clientId"].(string)
		exp := claims["exp"].(int64)

		if isInvalidToken(&clientId, &exp) {
			return c.String(http.StatusUnauthorized, "Invalid Token.")
		}
		return next(c)
	}
}

func ValidationHandler(target interface{}) error {
	config := &validator.Config{TagName: "validate"}
	validate := validator.New(config)
	if err := validate.Struct(target); err != nil {
		return err
	}

	return nil
}

func isInvalidToken(clientId *string, exp *int64) bool {
	client, err := s.GetClient(clientId)
	if err != nil {
		return false
	}

	return *client.UID == *clientId && *exp > time.Now().Unix()
}
