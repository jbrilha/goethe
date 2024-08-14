package handlers

import (
	"fmt"
	"goethe/db"
	"goethe/views/blog"
	"goethe/views/profile"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProfileBase(c echo.Context) error {
	username := c.Param("username")
    user, _ := db.GetUserAccount(username)
	return Render(c, profile.Index(user))
}

func Post(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
		fmt.Println("Invalid query param")
    }

    fmt.Println(id)
    post, _ := db.GetBlogPost(id)

	return Render(c, blog.Post(post))
}
