package handlers

import (
	"goethe/views/components"
	"goethe/views/routes"

	"github.com/labstack/echo/v4"
)

func Route404(c echo.Context) error {
	return Render(c, routes.Route404())
}

func NeedLogin(c echo.Context) error {
	c.Response().Header().Add("HX-Retarget", "#notifications")
	c.Response().Header().Add("HX-Reswap", "beforeend")

	return Render(c, components.Alert("Not logged in!", true))
}
