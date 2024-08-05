package data

type Book struct {
	Title  string
	Author string
}

func GetBooks() map[string][]Book {
	return map[string][]Book{
		"Books": {
			{Title: "Book of Disquiet", Author: "Fernando Pessoa"},
			{Title: "1984", Author: "George Orwell"},
			{Title: "My Life Had Stood a Loaded Gun", Author: "Emily Dickinson"},
		},
	}
}


