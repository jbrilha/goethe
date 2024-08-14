package handlers

import (
	"fmt"
	"goethe/data"
	"goethe/db"
	"goethe/views/blog"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
	return Render(c, blog.Index(data.GetPosts()))
}

func BlogPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		fmt.Println("Invalid query param")
    }

    fmt.Println(id)
    post := db.GetBlogPost(id)

	return Render(c, blog.Post(post))
}
