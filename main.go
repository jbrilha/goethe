package main

import (
	"goethe/config"
	"goethe/db"
	"goethe/env"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	env.Init()
	db.Init(env.DBConn())
	defer db.Close()

	config.ApplyEchoConfig(e)

	config.SetRoutes(e)

	e.Logger.Fatal(e.Start(env.Port()))
}
