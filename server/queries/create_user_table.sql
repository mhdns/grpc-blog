CREATE TABLE
IF NOT EXISTS user_table
(
    id SERIAL,
    email VARCHAR
(256),
    username VARCHAR
(256),
    pwd VARCHAR
(256),
    salt VARCHAR
(256),
PRIMARY KEY
(id)
);