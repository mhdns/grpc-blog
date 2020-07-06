CREATE TABLE IF NOT EXISTS blog(
    id SERIAL,
    title VARCHAR(256),
    create_date TIMESTAMP,
    post VARCHAR(256)
);