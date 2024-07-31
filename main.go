package main

import (
	"html/template"
	"io"

	"goethe/handlers"
	"goethe/config"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

const PORT = ":8000"

func main() {
	e := echo.New()

    config.ApplyEchoConfig(e)

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	e.GET("/", handlers.HandleBook)
	// e.POST("/add-book/", handlers.AddBook)

	e.Logger.Fatal(e.Start(PORT))
}
