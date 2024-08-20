package data

type Book struct {
	ID          int      `json:"id"`
	ISBN10      string   `json:"isbn10"`
	ISBN13      string   `json:"isbn13"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publishers  []string `json:"publishers"`
	Pages       int      `json:"pages"`
	PublishDate string   `json:"publish_date"`
	Description string   `json:"description"`
	Languages   []string `json:"languages"`
}

func GetBooks() []Book {
	return []Book{
		{Title: "Book of Disquiet", Authors: []string{"Fernando Pessoa"}},
		{Title: "1984", Authors: []string{"George Orwell"}},
		{Title: "My Life Had Stood a Loaded Gun", Authors: []string{"Emily Dickinson"}},
	}
}
