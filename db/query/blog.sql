-- name: CreateBlog :one
INSERT INTO blog (
  title,
  content,
  author_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetBlog :one
SELECT * FROM blog
WHERE id = $1 LIMIT 1;

-- name: ListBlog :many
SELECT b.id, title, content, name 
FROM blog as b
JOIN users as u
ON u.id = b.author_id
ORDER BY id;

-- name: UpdateBlogTitle :exec
UPDATE blog
SET title = $2
WHERE id = $1
RETURNING *;

-- name: UpdateBlogContent :exec
UPDATE blog
SET content = $2
WHERE id = $1
RETURNING *;

-- name: UpdateBlog :exec
UPDATE blog
SET title = $2,
content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blog
WHERE id = $1;