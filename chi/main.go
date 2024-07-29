package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Book struct {
	Title  string
	Author string
}

func Index(w http.ResponseWriter, r *http.Request) {
	books := map[string][]Book{
		"Books": {
			{Title: "Book of Disquiet", Author: "Fernando Pessoa"},
			{Title: "1984", Author: "George Orwell"},
			{Title: "My Life Had Stood a Loaded Gun", Author: "Emily Dickinson"},
		},
	}

	tmpl := template.Must(template.ParseFiles("../views/index.html"))
	tmpl.Execute(w, books)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")

	book := Book{Title: title, Author: author}

	tmpl := template.Must(template.ParseFiles("../views/index.html"))
	tmpl.ExecuteTemplate(w, "book-list-element", book)
}

const PORT = ":8081"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", Index)
	r.Post("/add-book/", addBook)

	log.Fatal(http.ListenAndServe(PORT, r))
}
