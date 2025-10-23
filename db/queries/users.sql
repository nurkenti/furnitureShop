-- name: CreateUser :one 
INSERT INTO users(
 id, email, password_hash,full_name, age, role
)VALUES( $1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: GetUserByID :one 
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;


-- name: UpdateUser :one
UPDATE users
SET 
  full_name = $2,
  age = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;


-- name: DeleteUserByEmail :exec
DELETE FROM users
WHERE email = $1;
