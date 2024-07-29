package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	colors "goethe/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Book struct {
	Title  string
	Author string
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Index(c echo.Context) error {
	books := map[string][]Book{
		"Books": {
			{Title: "Book of Disquiet", Author: "Fernando Pessoa"},
			{Title: "1984", Author: "George Orwell"},
			{Title: "My Life Had Stood a Loaded Gun", Author: "Emily Dickinson"},
		},
	}

	return c.Render(http.StatusOK, "index.html", books)
}

func addBook(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")

	book := Book{Title: title, Author: author}

	return c.Render(http.StatusOK, "book-list-element", book)
}

const PORT = ":8080"

func main() {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogProtocol:      true,
		LogLatency:       true,
		LogMethod:        true,
		LogStatus:        true,
		LogHost:          true,
		LogURI:           true,
		LogError:         true,
		LogContentLength: true,
		HandleError:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			dT := v.StartTime.Format(time.DateTime)
			method := fmt.Sprintf("%s%v%s", colors.BoldPurple, v.Method, colors.None)
			url := fmt.Sprintf("%s%v%v%s", colors.Cyan, v.Host, v.URI, colors.None)
			latColor := colors.Green
			if v.Latency > 10 * time.Second {
				latColor = colors.Red
			} else if v.Latency >= 5 * time.Second {
				latColor = colors.Orange
			} else if v.Latency >= 1*time.Second {
				latColor = colors.Yellow
			}
			lat := fmt.Sprintf("%s%v%s", latColor, v.Latency, colors.None)
			reqConLen := v.ContentLength
			if reqConLen == "" {
				reqConLen = "0"
			}
			reqConLen = fmt.Sprintf("%s%vB%s", colors.BoldBlue, reqConLen, colors.None)

			resConLen := strconv.Itoa(int(c.Response().Size))
			resConLen = fmt.Sprintf("%s%vB%s", colors.Blue, resConLen, colors.None)

			if v.Error == nil {
				log := "%v: %v %v [%v] — ꜛ%v ꜜ%v in %v\n"

				status := fmt.Sprintf("%s%v%s", colors.BoldGreen, v.Status, colors.None)
				fmt.Printf(log, dT, method, url, status, reqConLen, resConLen, lat)
			} else {
				log := "%v: %v %v [%v] — ꜛ%v ꜜ%v in %v\nError message: %v\n"

				status := fmt.Sprintf("%s%v%s", colors.BoldRed, v.Status, colors.None)
				error := strings.Split(v.Error.Error(), "message=")[1]
				fmt.Printf(log, dT, method, url, status, reqConLen, resConLen, lat, error)
			}

			return nil
		},
	}))

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

	e.GET("/", Index)
	e.POST("/add-book/", addBook)

	e.Logger.Fatal(e.Start(PORT))
}
