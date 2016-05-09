package controller
import (
    "log"
	"github.com/labstack/echo"
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

func ValidationHandler(path string) error {
    return nil
}
