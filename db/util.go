package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	// "goethe/data"
	// _ "github.com/lib/pq"
)

var db *pgxpool.Pool

func New(connStr string) {
	var err error
	// db, err = sql.Open("postgres", connStr)
	db, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(context.Background()); err != nil {
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

	err := db.QueryRow(context.Background(), query).Scan(&tz)
	if err != nil {
		log.Println(err)
	}

	log.Println(tz)
}

