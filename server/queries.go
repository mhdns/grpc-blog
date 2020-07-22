package main

const (
	createBlogTable = "CREATE TABLE IF NOT EXISTS blog (id SERIAL,title VARCHAR(256), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, post VARCHAR(256));"
	createUserTable = "CREATE TABLE IF NOT EXISTS user_table (id SERIAL, email VARCHAR(256), username VARCHAR(256), pwd VARCHAR(256), salt VARCHAR(256), PRIMARY KEY (id));"
	createBlog      = "INSERT INTO blog (title, post) VALUES ($1, $2) RETURNING id, title, created_at, post;"
	createUser      = "INSERT INTO user_table (email,username,pwd,salt) VALUES ($1, $2, $3, $4) RETURNING id, username, email;"
	getBlog         = "SELECT * FROM BLOG WHERE ID=$1;"
	getUserByID     = "SELECT * FROM USER_TABLE WHERE ID=$1;"
	getUserByEmail  = "SELECT * FROM USER_TABLE WHERE EMAIL=$1;"
	updateBlog      = "UPDATE blog SET title=$2, post=$3 WHERE id=$1 RETURNING id, title, created_at, post;"
	updateUser      = "UPDATE user_table SET email=$2, username=$3 WHERE id=$1 RETURNING id, username, email;"
	deleteBlog      = "DELETE FROM blog WHERE id=$1;"
	deleteUser      = "DELETE FROM user_table WHERE id=$1;"
)
