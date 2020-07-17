CREATE TABLE IF NOT EXISTS user (
    id SERIAL,
    email VARCHAR(256),
    name VARCHAR(256),
    password VARCHAR(256),
    salt VARCHAR(256),
    PRIMARY KEY id
);