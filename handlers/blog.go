package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	posts, err := db.GetBlogPosts()
	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}
	return Render(c, blog.Index(posts))
}

func PostSearch(c echo.Context) error {
	query := strings.TrimSpace(c.QueryParam("q"))
	if query == "" {
		return echo.ErrBadRequest
	}
	creator, fTerms, eTerms := parseQuery(query)
	p, err := db.SearchPosts(creator, fTerms, eTerms)

	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}

	if c.Request().Header.Get("HX-Request") == "" {
		// if it's not an htmx request it means it was a direct link access,
        // therefore I need to send @layouts.Base along with the results or else
        // it's just the results in plain html (no tailwind etc)
		return Render(c, blog.Index(p))
	}

	return Render(c, blog.Posts(p))
}

func parseQuery(query string) (string, []string, []string) {
	re := regexp.MustCompile(`"(.*?)"|from:(\S+)`)

	var creator string
	var exactTerms []string
	var fuzzyTerms []string

	matches := re.FindAllStringSubmatch(query, -1)
	for _, match := range matches {
		if match[2] != "" { // captured creator
			creator = match[2]
		} else if match[1] != "" { // captured string between quotes for exact matching
			exactTerms = append(exactTerms, match[1])
		}
	}

	for _, match := range matches {
		query = strings.ReplaceAll(query, match[0], "")
	}

	// query = strings.TrimSpace(query)
	if query != "" {
		fuzzyTerms = strings.Fields(query)
	}

	return creator, fuzzyTerms, exactTerms
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

	p, err := db.GetBlogPostByID(id)
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

	p, err := db.GetBlogPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": "Blog post not found"})
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, p)

}

func CreateBlogPostForm(c echo.Context) error {
	return Render(c, blog.CreatePost())
}

func validateTag(tag string) error {
	alphaNum := `^[a-zA-Z0-9_]+$`
	re := regexp.MustCompile(alphaNum)

	if !re.MatchString(tag) {
		return errors.New("Only alphanumeric characters and underscores in tags!")
	}

	return nil
}

func AddTag(c echo.Context) error {
	tag := c.QueryParam("tag")

	if err := validateTag(tag); err != nil {
		return alert(c, err.Error(), true)
	}

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
