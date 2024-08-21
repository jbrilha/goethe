package db

import (
	"database/sql"
	"fmt"
	"log"

	// "goethe/data"

	_ "github.com/lib/pq"
)

var db *sql.DB

func New(connStr string) {
	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createBookTable()
	createUserTable()
	createPostTable()

	// for _, post := range data.GetPosts() {
	// 	InsertBlogPost(post)
	// }

	// u := data.User{
	// 	Username:  "root",
	// 	Password:  "root",
	// 	Email:     "root@email.com",
	// 	CreatedAt: time.Now(),
	// }
	//
	// pk := InsertUser(u)
	// fmt.Println(pk)

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
		log.Fatal(err)
	}

	fmt.Println(tz)
}

