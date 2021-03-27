-- name: CreateComment :one
INSERT INTO user_comment (
  blog_id,
  user_id,
  coment
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListComment :many
SELECT * FROM user_comment
WHERE blog_id = $1;

-- -- name: UpdateCommentContent :one
-- UPDATE user_comment 
-- SET content = $2, 
-- WHERE id = $1
-- RETURNING *;

-- name: DeleteComment :exec
DELETE FROM user_comment
WHERE id = $1;