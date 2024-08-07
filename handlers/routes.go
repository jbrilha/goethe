package handlers

import (
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	e.Static("/public", "public")

	e.GET("/", HandleHome)

	e.GET("/bookshelf", BookshelfBase)
	e.POST("/bookshelf/add-book", AddBook)
	e.DELETE("/bookshelf/remove-book", RemoveBook)
	e.GET("/bookshelf/book", HandleBook)

	e.GET("/blog", BlogBase)
	e.GET("/blog/post", BlogPost)
}
