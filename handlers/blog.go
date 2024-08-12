package handlers

import (
	"fmt"
	"goethe/auth"
	"goethe/data"
	"goethe/util"
	"goethe/views/blog"
	"strconv"

	"github.com/labstack/echo/v4"
)

func BlogBase(c echo.Context) error {
	return Render(c, blog.Index(data.GetPosts()))
}

func BlogPost(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))
    if err != nil {
		fmt.Println("Invalid query param")
    }

    fmt.Println(id)
    post := data.GetPosts()[0]

    
    jwt, jerr := auth.CreateJWT("example")
    if jerr != nil {
		fmt.Println("Failed to create JWT", jerr)
    }

    err = util.WriteCookie(c, "JWT", jwt)
    if err != nil {
		fmt.Println("Cookie failed to write")
    }

	return Render(c, blog.Post(post))
}
