UPDATE user_table SET email=$2, name=$3 WHERE id=$1
RETURNING id, name;