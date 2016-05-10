package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/timakin/airshooter/config"
	"github.com/timakin/airshooter/route"
)

func main() {
	server := route.Init()
	server.Run(standard.New(config.Port))
}
