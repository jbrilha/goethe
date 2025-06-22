INSERT INTO users(username, email, password, created_at)
    VALUES($1, $2, $3, $4)
    RETURNING id;
