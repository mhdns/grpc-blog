INSERT INTO user_table
  (email,username,pwd,salt)
VALUES
  ($1, $2, $3, $4)
RETURNING id, username;