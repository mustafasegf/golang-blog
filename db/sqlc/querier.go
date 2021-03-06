// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateBlog(ctx context.Context, arg CreateBlogParams) (CreateBlogRow, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (CreateCommentRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteBlog(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	GetBlog(ctx context.Context, id int32) (GetBlogRow, error)
	GetComment(ctx context.Context, blogID int32) ([]GetCommentRow, error)
	GetOneComment(ctx context.Context, id int32) (GetOneCommentRow, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListBlog(ctx context.Context) ([]ListBlogRow, error)
	ListUser(ctx context.Context) ([]User, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) error
	UpdateBlogContent(ctx context.Context, arg UpdateBlogContentParams) error
	UpdateBlogTitle(ctx context.Context, arg UpdateBlogTitleParams) error
	UpdateComment(ctx context.Context, arg UpdateCommentParams) error
}

var _ Querier = (*Queries)(nil)
