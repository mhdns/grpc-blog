CREATE TABLE IF NOT EXISTS blog (
    id SERIAL,
    title VARCHAR(256),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    post VARCHAR(256)
);