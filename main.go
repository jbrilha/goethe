package main

import (
	"goethe/config"
	"goethe/db"
	"goethe/env"
	"goethe/util/policy"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	env.New()
    policy.New()
	db.New(env.DBConn())
	defer db.Close()

	config.ApplyEchoConfig(e)

	config.SetRoutes(e)

	e.Logger.Fatal(e.Start(env.Port()))
}
