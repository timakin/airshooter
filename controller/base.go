package controller

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v8"
	"log"
)

func ValidateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("validate")
		return next(c)
	}
}

func AuthClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("JWT Auth")
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
