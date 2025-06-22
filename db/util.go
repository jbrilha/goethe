package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func New(connStr string) {
	ctx := context.Background()
	var err error
	db, err = pgxpool.New(ctx, connStr)

	if err != nil {
		log.Println(err)
	}

	if err = db.Ping(ctx); err != nil {
		log.Println(err)
	}

	createBookTable()
	createUserTable()
	createPostsTable()
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

