package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/timakin/airshooter/config"
	"github.com/timakin/airshooter/route"
)

func main() {
	server := route.Init()
	settings := config.Load()
	server.Run(standard.New(settings.AppPort))
}
