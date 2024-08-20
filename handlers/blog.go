package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"goethe/db"
	"goethe/views/blog"
	"goethe/views/routes"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
	return Render(c, blog.Index(db.GetBlogPosts()))
}

func BlogPost(c echo.Context) error {
	p := c.Param("id")

	if idStr := strings.TrimSuffix(p, ".json"); idStr != p {
		id, err := strconv.Atoi(idStr)

		if err != nil {
			fmt.Println("Invalid param:", idStr)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid blog post ID"})
		}

		post, err := db.GetBlogPost(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Blog post not found"})
		}

		return c.JSON(http.StatusOK, post)
	}

	id, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Invalid param")
	}

	post, err := db.GetBlogPost(id)
	if err != nil {
		fmt.Println(err)
		return Render(c, routes.Route404())
	}

	return Render(c, blog.Post(post))
}
