package db

import (
	"database/sql"
	"log"

	// "goethe/data"

	_ "github.com/lib/pq"
)

var db *sql.DB

func New(connStr string) {
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	createBookTable()
	createUserTable()
	createPostTable()

	// for _, post := range data.GetPosts() {
	// 	InsertBlogPost(&post)
	// }
}

func emptyNullString(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }

    return ""
}

func Close() {
	db.Close()

	log.Println("DB closed")
}

func checkTz() {
	query := `SHOW timezone`

	var tz string

	err := db.QueryRow(query).Scan(&tz)
	if err != nil {
		log.Println(err)
	}

	log.Println(tz)
}

