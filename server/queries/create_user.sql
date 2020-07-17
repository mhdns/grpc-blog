INSERT INTO user (email,name,password,salt) VALUES ($1, $2, $3, $4)
RETURNING id, name;