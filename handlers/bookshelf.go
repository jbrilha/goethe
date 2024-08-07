package handlers

import (
	"goethe/data"
	"goethe/views/bookshelf"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) error {
	title := c.FormValue("book-title")
	author := c.FormValue("book-author")

	book := data.Book{Title: title, Author: author}

	return Render(c, bookshelf.AddBook(book))
}

func RemoveBook(c echo.Context) error {
	title := c.FormValue("book-title")
	author := c.FormValue("book-author")

	book := data.Book{Title: title, Author: author}

	return Render(c, bookshelf.RemoveBook(book))
}

func HandleBook(c echo.Context) error {
    book := data.Book {
        Title: "AA",
        Author: "BB",
    }
    return Render(c, bookshelf.Show(book))
}

func BookshelfBase(c echo.Context) error {
    return Render(c, bookshelf.Index(data.GetBooks()))
}
