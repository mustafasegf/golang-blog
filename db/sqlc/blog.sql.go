// Code generated by sqlc. DO NOT EDIT.
// source: blog.sql

package db

import (
	"context"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blog (
  title,
  content,
  author_id
) VALUES (
  $1, $2, $3
) RETURNING id, title, content, author_id
`

type CreateBlogParams struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int32  `json:"author_id"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog, arg.Title, arg.Content, arg.AuthorID)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blog
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT b.id, title, content, name, u.id as userid 
FROM blog as b
JOIN users as u
ON u.id = b.author_id
WHERE b.id = $1 LIMIT 1
`

type GetBlogRow struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Userid  int32  `json:"userid"`
}

func (q *Queries) GetBlog(ctx context.Context, id int32) (GetBlogRow, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i GetBlogRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Name,
		&i.Userid,
	)
	return i, err
}

const listBlog = `-- name: ListBlog :many
SELECT b.id, title, content, name, u.id as userid
FROM blog as b
JOIN users as u
ON u.id = b.author_id
ORDER BY id
`

type ListBlogRow struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Name    string `json:"name"`
	Userid  int32  `json:"userid"`
}

func (q *Queries) ListBlog(ctx context.Context) ([]ListBlogRow, error) {
	rows, err := q.db.QueryContext(ctx, listBlog)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListBlogRow
	for rows.Next() {
		var i ListBlogRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Name,
			&i.Userid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :exec
UPDATE blog
SET title = $2,
content = $3
WHERE id = $1
RETURNING id, title, content, author_id
`

type UpdateBlogParams struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) error {
	_, err := q.db.ExecContext(ctx, updateBlog, arg.ID, arg.Title, arg.Content)
	return err
}

const updateBlogContent = `-- name: UpdateBlogContent :exec
UPDATE blog
SET content = $2
WHERE id = $1
RETURNING id, title, content, author_id
`

type UpdateBlogContentParams struct {
	ID      int32  `json:"id"`
	Content string `json:"content"`
}

func (q *Queries) UpdateBlogContent(ctx context.Context, arg UpdateBlogContentParams) error {
	_, err := q.db.ExecContext(ctx, updateBlogContent, arg.ID, arg.Content)
	return err
}

const updateBlogTitle = `-- name: UpdateBlogTitle :exec
UPDATE blog
SET title = $2
WHERE id = $1
RETURNING id, title, content, author_id
`

type UpdateBlogTitleParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateBlogTitle(ctx context.Context, arg UpdateBlogTitleParams) error {
	_, err := q.db.ExecContext(ctx, updateBlogTitle, arg.ID, arg.Title)
	return err
}
