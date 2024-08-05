package handlers

import (
	"goethe/views/blog"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
    return Render(c, blog.Index())
}
