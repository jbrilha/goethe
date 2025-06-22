SELECT * FROM posts WHERE $1 = ANY(tags) ORDER BY created_at DESC;
