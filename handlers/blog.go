package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"goethe/auth"
	"goethe/data"
	"goethe/db"
	"goethe/util"
	"goethe/views/blog"
	"goethe/views/routes"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
    fmt.Printf("c.QueryParam(\"c\"): %v\n", c.QueryParam("c"))
    timestamp, err := time.Parse("01-02-2006 15:04:05:00", c.QueryParam("ts"))
	if err != nil {
		log.Println("ts", err)
	}

	posts, err := db.GetBlogPosts(id, timestamp)
	if err != nil {
		log.Println(err)
		return Render(c, routes.Route404())
	}

	if timestamp.IsZero() {
		return Render(c, blog.Index(posts))
	}
	return Render(c, blog.Posts(posts))
}

func PostSearch(c echo.Context) error {
	query := strings.TrimSpace(c.QueryParam("q"))

	sp := parseQuery(query)
	p, err := db.SearchPosts(sp)

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

func parseQuery(query string) db.PostSearchParams {
	re := regexp.MustCompile(`"(.*?)"|from:(\S+)|#(\w+)`)

	sp := db.PostSearchParams{}

	// var creator string
	// var exactTerms []string
	// var fuzzyTerms []string
	// var tags []string

	matches := re.FindAllStringSubmatch(query, -1)
	for _, match := range matches {
		if match[2] != "" { // captured creator
			sp.Creator = match[2]
		} else if match[1] != "" { // captured string between quotes for exact matching
			sp.ExactTerms = append(sp.ExactTerms, match[1])
		} else if match[3] != "" { // captured tags
			sp.Tags = append(sp.Tags, match[3])
		}
	}

	for _, match := range matches {
		query = strings.ReplaceAll(query, match[0], "")
	}

	// query = strings.TrimSpace(query)
	if query != "" {
		sp.FuzzyTerms = strings.Fields(query)
	}

	return sp
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

	tagSlice := strings.Fields(tags)
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
