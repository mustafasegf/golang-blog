-- name: CreateBlog :one
INSERT INTO blogs (
  title,
  content,
  author_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetBlog :one
SELECT b.id, title, content, name, u.id as userid 
FROM blogs as b
JOIN users as u
ON u.id = b.author_id
WHERE b.id = $1 LIMIT 1;

-- name: ListBlog :many
SELECT b.id, title, content, name, u.id as userid
FROM blogs as b
JOIN users as u
ON u.id = b.author_id
ORDER BY id;

-- name: UpdateBlogTitle :exec
UPDATE blogs
SET title = $2
WHERE id = $1
RETURNING *;

-- name: UpdateBlogContent :exec
UPDATE blogs
SET content = $2
WHERE id = $1
RETURNING *;

-- name: UpdateBlog :exec
UPDATE blogs
SET title = $2,
content = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;