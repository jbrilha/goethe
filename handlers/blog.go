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

func PostsByTag(c echo.Context) error {
	tag := c.Param("tag")

	// if idStr := strings.TrimSuffix(param, ".json"); idStr != param {
	// 	return BlogPostJSON(c, idStr)
	// }

	p, err := db.GetBlogPostsByTag(tag)
	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}

	return Render(c, blog.Index(p))
}

func CreatorCard(c echo.Context) error {
    username := c.Param("creator")

	u, err := db.GetUserAccountByUsername(username)
	if err != nil {
		log.Println(err)
        return Render(c, routes.Route404())
	}
    log.Println(u)

    return Render(c, blog.CreatorCard(u))
}

func BlogPost(c echo.Context) error {
	param := c.Param("id")

	if idStr := strings.TrimSuffix(param, ".json"); idStr != param {
		return BlogPostJSON(c, idStr)
	}

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Invalid param")
	}

	p, err := db.GetBlogPost(id)
	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}

	go func(id int) {
		err = db.IncrPostViews(id)

		if err != nil {
			log.Println("err in incrPostViews goroutine", err)
		}
	}(id)

    p.Views += 1 // just to reflect current visit on page
	return Render(c, blog.Post(p))
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

	p, err := db.GetBlogPost(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Blog post not found"})
	}

	return c.JSON(http.StatusOK, p)

}

func CreateBlogPostForm(c echo.Context) error {
	return Render(c, blog.CreatePost())
}

func AddTag(c echo.Context) error {
    tag := c.FormValue("new-tag")

	return Render(c, blog.Tag(tag))
}

func CreateBlogPostSubmission(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
    tags := c.FormValue("tags")

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

    tagSlice := strings.Split(tags, ",")
    p := data.Post{Creator: creator, Title: title, Tags: tagSlice, Content: content}

	_, err = db.InsertBlogPost(&p)
	if err != nil {
		log.Println("err in insertion:", err)
	}
	log.Println(p.ID)

    p.Views += 1
	c.Response().Header().Add("HX-Push-Url", fmt.Sprintf("/posts/%v", p.ID))
	return Render(c, blog.Post(p))
}
