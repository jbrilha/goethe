package handlers

import (
	"fmt"
	"goethe/data"
	"goethe/views/blog"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
	return Render(c, blog.Index(data.GetPosts()))
}

func BlogPost(c echo.Context) error {
	id, e := strconv.Atoi(c.QueryParam("id"))
    if e != nil {
		fmt.Println("Invalid query param")
    }

	return Render(c, blog.Post(data.GetPosts()[id]))
}
