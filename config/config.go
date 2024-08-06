package config

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"goethe/util/ansi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/joho/godotenv"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func Port() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong loading .env file;\nDefaulting to port :8080")
		return ":8080"
	}

	return os.Getenv("PORT")
}

func ApplyEchoConfig(e *echo.Echo) {

	t := &Template{
		Templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t

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
			method := fmt.Sprintf("%s%v%s", ansi.BoldPurple, v.Method, ansi.None)
			url := fmt.Sprintf("%s%v%v%s", ansi.Cyan, v.Host, v.URI, ansi.None)

			latColor := ansi.Green
			if v.Latency > 10*time.Second {
				latColor = ansi.Red
			} else if v.Latency >= 5*time.Second {
				latColor = ansi.Orange
			} else if v.Latency >= 1*time.Second {
				latColor = ansi.Yellow
			}
			lat := fmt.Sprintf("%s%v%s", latColor, v.Latency, ansi.None)

			reqConLen := v.ContentLength
			if reqConLen == "" {
				reqConLen = "0"
			}
			reqConLen = fmt.Sprintf("%s%vB%s", ansi.BoldBlue, reqConLen, ansi.None)

			resConLen := strconv.Itoa(int(c.Response().Size))
			resConLen = fmt.Sprintf("%s%vB%s", ansi.Blue, resConLen, ansi.None)

			if v.Error == nil {
				log := "%v: %v %v [%v] — ꜛ%v ꜜ%v in %v\n"

				status := fmt.Sprintf("%s%v%s", ansi.BoldGreen, v.Status, ansi.None)

				fmt.Printf(log, dT, method, url, status, reqConLen, resConLen, lat)
			} else {
				log := "%v: %v %v [%v] — ꜛ%v ꜜ%v in %v\nError message: %v\n"

				status := fmt.Sprintf("%s%v%s", ansi.BoldRed, v.Status, ansi.None)
				error := strings.Split(v.Error.Error(), "message=")[1]

				fmt.Printf(log, dT, method, url, status, reqConLen, resConLen, lat, error)
			}

			return nil
		},
	}))
}
