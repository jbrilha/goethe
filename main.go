package main

import (
	"goethe/config"
	"goethe/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ApplyEchoConfig(e)

	handlers.SetRoutes(e)

	e.Logger.Fatal(e.Start(config.Port()))
}
