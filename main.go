package main

import (
	"html/template"

	"goethe/config"
	"goethe/handlers"

	"github.com/labstack/echo/v4"
)

const PORT = ":8000"

func main() {
	e := echo.New()

	config.ApplyEchoConfig(e)

	t := &config.Template{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	handlers.SetRoutes(e)

	e.Logger.Fatal(e.Start(PORT))
}
