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
SELECT * FROM blog
ORDER BY TITLE;

-- -- name: UpdateBlogContent :one
-- UPDATE blog 
-- SET content = $2, 
-- WHERE id = $1
-- RETURNING *;

-- -- name: UpdateBlogTitle :one
-- UPDATE blog 
-- SET title = $2, 
-- WHERE id = $1
-- RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blog
WHERE id = $1;