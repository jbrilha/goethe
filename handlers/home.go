package handlers

import (
	"goethe/views/home"

	"github.com/labstack/echo/v4"
)

func HandleHome(c echo.Context) error {
	return Render(c, home.Index())
}
