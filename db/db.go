package db

import (
	"database/sql"
	"fmt"
	"goethe/data"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(connStr string) {
	var err error // instead of := bc I need to initialize the var db
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createBlogTable()
    checkTz()
	post := data.GetPosts()[2]
	insertBlogPost(post)
}

func Close() {
	db.Close()

	log.Println("DB closed")
}

func createBlogTable() {
	db.Exec("DROP TABLE blog")
	query := `CREATE TABLE IF NOT EXISTS blog(
                id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                title VARCHAR NOT NULL,
                content TEXT NOT NULL,
                createdAt timestamp
    )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func checkTz() {
	query := `SHOW timezone`

	var tz string

	err := db.QueryRow(query).Scan(&tz)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println(tz)
}

func insertBlogPost(p data.Post) int {
	query := `INSERT INTO blog(title, content, createdAt)
                VALUES($1, $2, $3)
                RETURNING id`

	var pk int

	err := db.QueryRow(query, p.Title, p.Content, time.Now()).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk
}

func getBlogPost(id int) {
	query := `SELECT title, content FROM blog where ID = $1`

	var title string
	var content string

	err := db.QueryRow(query, id).Scan(&title, &content)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("nooooooooo %d", id)
		}
		log.Fatal(err)
	}

	log.Println(title)
	log.Println(content)
}

func GetBlogPosts() []data.Post {
	query := `SELECT * FROM blog`

	var id int
	var title string
	var content string
	var timestamp time.Time

	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("nooooooooo roooooooooooows")
		}
		log.Fatal(err)
	}

	defer rows.Close()

	posts := []data.Post{}

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &timestamp)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, data.Post{
			ID:        id,
			Title:     title,
			Content:   content,
			Timestamp: timestamp,
		})
	}

	return posts
}