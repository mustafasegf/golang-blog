-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  name
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY name;