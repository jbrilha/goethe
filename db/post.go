package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"goethe/data"
)

type PostSearchParams struct {
	Creator    string
	Tags       []string
	FuzzyTerms []string
	ExactTerms []string
	ID         int
	Timestamp  time.Time
	Limit      int
    Refresh bool
}

func InsertBlogPost(p *data.Post) (int, error) {
	tx, err := db.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return 0, err
	}

	defer tx.Rollback(context.Background())

	query := `INSERT INTO post(creator, title, tags, content, created_at)
                VALUES($1, $2, $3, $4, $5)
                RETURNING id`

	err = tx.QueryRow(
        context.Background(),
		query,
		p.Creator,
		p.Title,
		p.Tags,
		p.Content,
		time.Now(),
	).Scan(
		&p.ID,
	)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if err = tx.Commit(context.Background()); err != nil {
		log.Println(err)
		return 0, err
	}

	return p.ID, nil
}

func IncrPostViews(id int) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}

	defer tx.Rollback(context.Background())

	query := `UPDATE post SET views = views + 1 WHERE id = $1`

	_, err = tx.Exec(context.Background(), query, id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("nooooooooo %d", id)
			return err
		}
		log.Println("wtf", id)
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetBlogPostByID(id int) (data.Post, error) {
	query := `SELECT * FROM post WHERE id = $1`

	post := data.Post{}

	err := db.QueryRow(context.Background(), query, id).Scan(
		&post.ID,
		&post.Creator,
		&post.Title,
		&post.Tags,
		&post.Content,
		&post.Views,
		&post.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("nooooooooo %d", id)
			return data.Post{}, err
		}
		log.Println("wtf", id)
		return data.Post{}, err
	}

	return post, nil
}

func GetBlogPosts(id int, timestamp time.Time) ([]data.Post, error) {
	query := `SELECT * FROM post
            WHERE (id > $1 AND created_at < $2) OR (created_at < $2)
            ORDER BY (created_at, id) DESC
            LIMIT 20`

	return getPosts(query, id, timestamp)
}

func SearchPosts(sp PostSearchParams) ([]data.Post, error) {
	var query strings.Builder
	args := []any{}
	offset := 2
	fCount := len(sp.FuzzyTerms)
	eCount := len(sp.ExactTerms)
	tCount := len(sp.Tags)

	query.WriteString(`SELECT * FROM post WHERE `)
    if sp.Creator == "" {
        if sp.Refresh {
		    query.WriteString(`(created_at > $1)`)
        } else {
		    query.WriteString(`(created_at < $1)`)
        }
	} else {
		query.WriteString(`creator = $1 AND `)
        if sp.Refresh {
		    query.WriteString(`(created_at > $2)`)
        } else {
		    query.WriteString(`(created_at < $2)`)
        }
		offset = 3

		args = []any{sp.Creator}
	}
	    args = append(args, sp.Timestamp)

	if tCount > 0 || fCount > 0 || eCount > 0 {
		query.WriteString(` AND (`)
	}

	q := []string{}

	if fCount > 0 {
		query.WriteString(`(`)
		for _, term := range sp.FuzzyTerms { // TODO already looping through terms in the handler, figure out optimization
			q = append(
				q,
				`(content ILIKE '%' || $`+fmt.Sprint(offset)+
					` || '%' OR title ILIKE '%' || $`+fmt.Sprint(offset)+` || '%')`)
			args = append(args, term)

			offset += 1
		}
		query.WriteString(strings.Join(q, ` OR `))
		query.WriteString(`)`)
	}

	if eCount > 0 {
		q = []string{}
		if fCount > 0 {
			query.WriteString(" AND (")
		}

		for _, term := range sp.ExactTerms {
			q = append(
				q,
				`(content ~* ('\y' || $`+fmt.Sprint(offset)+
					` || '\y') OR title ~* ('\y' || $`+fmt.Sprint(offset)+` || '\y'))`)
			args = append(args, term)

			offset += 1
		}
		query.WriteString(strings.Join(q, " OR "))

		if fCount > 0 {
			query.WriteString(")")
		}
	}

	if tCount > 0 {
		if fCount > 0 || eCount > 0 {
			query.WriteString(" AND (")
		}

		query.WriteString("tags && $" + fmt.Sprint(offset))
		args = append(args, sp.Tags)

		if fCount > 0 || eCount > 0 {
			query.WriteString(")")
		}
	}

	if tCount > 0 || fCount > 0 || eCount > 0 {
		query.WriteString(`)`)
	}

	args = append(args, sp.Limit)
	query.WriteString(` ORDER BY created_at DESC LIMIT $` + fmt.Sprint(offset))

	return getPosts(query.String(), args...)
}

func SearchPostsByCreator(creator string) ([]data.Post, error) {
	query := `SELECT * FROM post WHERE creator ILIKE '%' || $1 || '%' ORDER BY created_at DESC`

	return getPosts(query, creator)
}

func SearchPostsByTag(tag string) ([]data.Post, error) {
	query := `SELECT * FROM post WHERE $1 = ANY(tags) ORDER BY created_at DESC`

	return getPosts(query, tag)
}

func getPosts(query string, args ...any) ([]data.Post, error) {
	rows, err := db.Query(context.Background(), query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("nooooooooo roooooooooooows")
		}
		log.Println(err)
	}

	defer rows.Close()

	posts := []data.Post{}

	for rows.Next() {
		var post data.Post
		err := rows.Scan(
			&post.ID,
			&post.Creator,
			&post.Title,
			&post.Tags,
			&post.Content,
			&post.Views,
			&post.CreatedAt,
		)
		if err != nil {
			log.Println(err)
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func createPostTable() {
	// db.Exec("DROP TABLE post")
	query := `CREATE TABLE IF NOT EXISTS post(
                id int PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                creator VARCHAR NOT NULL,
                title VARCHAR NOT NULL,
                tags VARCHAR ARRAY DEFAULT NULL,
                content TEXT NOT NULL,
                views INT DEFAULT 0,
                created_at timestamp NOT NULL
    )`

	_, err := db.Exec(context.Background(), query)
	if err != nil {
		log.Println(err)
	}
}
