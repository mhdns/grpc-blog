UPDATE user_table SET email=$2, username=$3 WHERE id=$1
RETURNING id, username;