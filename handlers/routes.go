package handlers

import (
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
    e.Static("/public", "public")

	e.GET("/", HandleHome)

	e.GET("/books", BooksBase)
	e.GET("/blog", BlogBase)

	e.GET("/book", HandleBook)
	e.POST("/add-book", AddBook)
	e.DELETE("/remove-book", RemoveBook)
}
