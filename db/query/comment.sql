-- name: CreateComment :one
INSERT INTO comments (
  blog_id,
  user_id,
  comment
) VALUES (

  $1, $2, $3
) RETURNING *, (SELECT u.name from users as u WHERE u.id = $2) AS name;

-- name: GetComment :many
SELECT *, (SELECT u.name from users as u WHERE u.id = c.user_id) AS name 
FROM comments as c
WHERE c.blog_id = $1;

-- name: GetOneComment :one
SELECT *, (SELECT u.name from users as u WHERE u.id = c.user_id) AS name 
FROM comments as c
WHERE c.id = $1
LIMIT 1;

-- name: UpdateComment :exec
UPDATE comments AS c
SET comment = $2
WHERE c.id = $1
RETURNING *, (SELECT u.name from users as u WHERE u.id = c.user_id) AS name ;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;