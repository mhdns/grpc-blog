INSERT INTO blog (title, post) VALUES ($1, $2)
RETURNING id, title, created_at, post;