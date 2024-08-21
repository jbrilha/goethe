package handlers

import (
	"context"

	"goethe/util"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, comp templ.Component) error {
	ctx := c.Request().Context()

	cookie, err := util.ReadCookie(c, "JWT")
	if err == nil {
		ctx = context.WithValue(context.Background(), "JWT", cookie.Value)
		ctx = c.Request().WithContext(ctx).Context()
	}

	return comp.Render(ctx, c.Response())
}
