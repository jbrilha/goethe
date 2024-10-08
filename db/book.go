package db

import (
	"database/sql"
	"log"

	"goethe/data"

	"github.com/lib/pq"
)

func InsertBook(b *data.Book) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	defer tx.Rollback()

	query := `INSERT INTO book(isbn10, isbn13, title, authors, publishers, publish_date, pages, description, languages)
                VALUES(NULLIF($1, ''), $2, $3, $4, $5, $6, $7, $8, $9)
                RETURNING id`

	err = db.QueryRow(
		query,
        b.ISBN10,
		b.ISBN13,
		b.Title,
		pq.Array(b.Authors),
		pq.Array(b.Publishers),
		b.PublishDate,
		b.Pages,
		b.Description,
		pq.Array(b.Languages),
	).Scan(&b.ID)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		log.Println(err)
		return 0, err
	}

	return b.ID, nil
}

func GetBookByISBN(isbn string) (data.Book, error) {
	query := `SELECT * FROM book WHERE username = $1`

	b := data.Book{}

    var isbn10 sql.NullString

	err := db.QueryRow(
		query, isbn).Scan(
		&b.ID,
		&isbn10,
		&b.ISBN13,
		&b.Title,
		pq.Array(&b.Authors),
		pq.Array(&b.Publishers),
		&b.PublishDate,
		&b.Pages,
		&b.Description,
		pq.Array(&b.Languages),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return data.Book{}, err
		}
		log.Println("other err:", err)
		return data.Book{}, err
	}

    b.ISBN10 = emptyNullString(isbn10)
	return b, nil
}

func GetBooks() []data.Book {
	query := `SELECT * FROM book`

	var (
		id          int
		isbn10      sql.NullString
		isbn13      string
		title       string
		authors     []string
		publishers  []string
		publishDate string
		pages       int
		description string
		languages   []string
	)

	rows, err := db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("nooooooooo roooooooooooows")
		}
		log.Println(err)
	}

	defer rows.Close()

	books := []data.Book{}

	for rows.Next() {
		err := rows.Scan(
			&id,
			&isbn10,
			&isbn13,
			&title,
			pq.Array(&authors),
			pq.Array(&publishers),
			&publishDate,
			&pages,
			&description,
			pq.Array(&languages),
		)
		if err != nil {
			log.Println(err)
		}

		books = append(books, data.Book{
			ID:          id,
			ISBN10:      emptyNullString(isbn10),
			ISBN13:      isbn13,
			Title:       title,
			Authors:     authors,
			Publishers:  publishers,
			PublishDate: publishDate,
			Pages:       pages,
			Description: description,
			Languages:   languages,
		})
	}

	return books
}

func createBookTable() {
	// db.Exec("DROP TABLE book")
	query := `CREATE TABLE IF NOT EXISTS book(
                id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                isbn10 CHAR(10),
                isbn13 CHAR(13) NOT NULL UNIQUE,
                title VARCHAR NOT NULL,
                authors VARCHAR ARRAY NOT NULL,
                publishers VARCHAR ARRAY,
                publish_date VARCHAR,
                pages int NOT NULL,
                description TEXT,
                languages VARCHAR ARRAY
    )`

	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
	}

	// query = `CREATE UNIQUE INDEX unq_isbn10 ON book (isbn10) WHERE isbn10 IS NOT NULL;`
	//
	// _, err = db.Exec(query)
	// if err != nil {
	// 	log.Println(err)
	// }
}
