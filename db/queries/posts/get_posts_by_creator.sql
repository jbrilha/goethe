SELECT * FROM posts WHERE creator ILIKE '%' || $1 || '%' ORDER BY created_at DESC;
