-- name: CreateComment :one
INSERT INTO comments (
  blog_id,
  user_id,
  comment
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetComment :many
SELECT * FROM comments
WHERE blog_id = $1;

-- name: UpdateComment :exec
UPDATE comments
SET comment = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;