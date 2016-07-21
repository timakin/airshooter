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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": clientId,
		"iss": "airshooter",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	// Encode and sign a token and send it as a response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	tokenRecord := m.AccessToken{Token: &t}
	accessToken, err := s.InsertToken(&tokenRecord)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

	return echo.ErrUnauthorized
}
