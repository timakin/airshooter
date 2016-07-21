package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	m "github.com/timakin/airshooter/model"
	s "github.com/timakin/airshooter/service"
)

func Authenticate(c echo.Context) error {
	// DBのclientsテーブルに保存されたclient_id & passwordと照合して、
	// 正規のクライアントであればtokenを引き渡す
	clientId := c.FormValue("client_id")
	clientSecret := c.FormValue("client_secret")

	client, err := s.GetClient(&clientId)
	if err != nil {
		return err
	}

	if *client.Secret != clientSecret {
		return echo.ErrUnauthorized
	}

	// Generate token
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": clientId,
		"iss": "airshooter",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	// Encode and sign a token and send it as a response.
	signed, err := payload.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	token := m.AccessToken{Token: &signed}
	accessToken, err := s.InsertToken(&token)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": *accessToken.Token,
	})
}
