package handlers

import (
	"goethe/data"
	"goethe/views/books"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) error {
	title := c.FormValue("book-title")
	author := c.FormValue("book-author")

	book := data.Book{Title: title, Author: author}

	return Render(c, books.AddBook(book))
}

func RemoveBook(c echo.Context) error {
	title := c.FormValue("book-title")
	author := c.FormValue("book-author")

	book := data.Book{Title: title, Author: author}

	return Render(c, books.RemoveBook(book))
}

func HandleBook(c echo.Context) error {
    book := data.Book {
        Title: "AA",
        Author: "BB",
    }
    return Render(c, books.Show(book))
}

func BooksBase(c echo.Context) error {
    return Render(c, books.Index(data.GetBooks()))
}
