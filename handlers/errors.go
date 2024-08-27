package handlers

import (
	"goethe/views/routes"

	"github.com/labstack/echo/v4"
)

func Route404(c echo.Context) error {
	return Render(c, routes.Route404())
}

func NeedLogin(c echo.Context) error {
    return alert(c, "Not logged in!", true)
}
