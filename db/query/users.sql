-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  role,
  created,
  updated
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM users
ORDER BY name;

-- -- name: UpdateUser :one
-- UPDATE users 
-- SET username = $2, 
-- updated = NOW()
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteUser :exec
-- DELETE FROM users
-- WHERE id = $1;