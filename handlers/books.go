package handlers

import (
	"goethe/views/books"
	"goethe/data"

	"github.com/labstack/echo/v4"
)

// func getBooks() map[string][]Book {
// 	return map[string][]Book{
// 		"Books": {
// 			{Title: "Book of Disquiet", Author: "Fernando Pessoa"},
// 			{Title: "1984", Author: "George Orwell"},
// 			{Title: "My Life Had Stood a Loaded Gun", Author: "Emily Dickinson"},
// 		},
// 	}
// }
//
// func AddBook(c echo.Context) error {
// 	title := c.FormValue("title")
// 	author := c.FormValue("author")
//
// 	book := Book{Title: title, Author: author}
//
// 	return c.Render(http.StatusOK, "book-list-element", book)
// }

func HandleBook(c echo.Context) error {
    book := data.Book {
        Title: "AA",
        Author: "BB",
    }
    return Render(c, books.Show(book))
}
