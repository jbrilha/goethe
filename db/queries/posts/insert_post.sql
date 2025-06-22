INSERT INTO posts(creator, title, tags, content, created_at)
    VALUES($1, $2, $3, $4, $5)
    RETURNING id;
