package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"goethe/auth"
	"goethe/data"
	"goethe/db"
	"goethe/util"
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
		return BlogPostJSON(c, idStr)
	}

	id, err := strconv.Atoi(p)
	if err != nil {
		log.Println("Invalid param")
	}

	post, err := db.GetBlogPost(id)
	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}

	err = db.IncrPostViews(id)
	if err != nil {
		log.Println(err)
	}

	return Render(c, blog.Post(post))
}

func BlogPostJSON(c echo.Context, idStr string) error {
	id, err := strconv.Atoi(idStr)

	if err != nil {
		log.Println("Invalid param:", idStr)
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": "Invalid post ID — should be a digit"},
		)
	}

	post, err := db.GetBlogPost(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Blog post not found"})
	}

	return c.JSON(http.StatusOK, post)

}

func CreateBlogPostForm(c echo.Context) error {
	return Render(c, blog.CreatePost())
}

func CreateBlogPostSubmission(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	jwtCookie, err := util.ReadCookie(c, "JWT")
	if err != nil {
		return c.JSON(http.StatusBadRequest, data.Post{})
	}
	token, err := auth.ValidateJWT(jwtCookie.Value)
	if err != nil {
		return c.JSON(http.StatusForbidden, data.Post{})
	}
	creator, err := token.Claims.GetSubject()
	if err != nil {
		return c.JSON(http.StatusTeapot, data.Post{})
	}

	p := data.Post{Creator: creator, Title: title, Content: content}

    _, err = db.InsertBlogPost(&p)
	if err != nil {
		log.Println("err in insertion:", err)
	}
	log.Println(p.ID)


	c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/posts/%v", p.ID))
	return Render(c, blog.Post(p))
}
