package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"goethe/data"

	_ "github.com/lib/pq"
)

func InsertBlogPost(p data.Post) int {
	query := `INSERT INTO post(creator, title, content, created_at)
                VALUES($1, $2, $3, $4)
                RETURNING id`

	var pk int

	err := db.QueryRow(query, p.Creator, p.Title, p.Content, time.Now()).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}

	return pk
}

func GetBlogPost(id int) (data.Post, error) {
	query := `SELECT * FROM post WHERE id = $1`

	post := data.Post{}

	err := db.QueryRow(query, id).Scan(
		&post.ID,
		&post.Creator,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.Views,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("nooooooooo %d", id)
			return data.Post{}, err
		}
		fmt.Println("wtf", id)
		return data.Post{}, err
	}

	return post, nil
}

func GetBlogPosts() []data.Post {
	query := `SELECT * FROM post`

	var id int
	var title string
	var creator string
	var content string
	var views int
	var createdAt time.Time

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
		err := rows.Scan(&id, &creator, &title, &content, &createdAt, &views)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, data.Post{
			ID:        id,
			Creator:   creator,
			Title:     title,
			Content:   content,
			CreatedAt: createdAt,
			Views:     views,
		})
	}

	return posts
}

func createPostTable() {
	// db.Exec("DROP TABLE post")
	query := `CREATE TABLE IF NOT EXISTS post(
                id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                creator VARCHAR NOT NULL,
                title VARCHAR NOT NULL,
                content TEXT NOT NULL,
                views INT DEFAULT 0,
                created_at timestamp NOT NULL
    )` // TODO drop table before using views

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
