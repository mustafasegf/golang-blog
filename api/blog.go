package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
)

type createBlogRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	AuthorID int32  `json:"authorId" binding:"required"`
}

type blogResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func newBlogResponse(blog db.Blog) blogResponse {
	return blogResponse{
		Title:   blog.Title,
		Content: blog.Content,
	}
}

func (server *Server) createBlog(ctx *gin.Context) {
	var req createBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBlogParams{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: req.AuthorID,
	}

	blog, err := server.store.CreateBlog(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newBlogResponse(blog)
	ctx.JSON(http.StatusOK, rsp)
}

type getBlogRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getBlog(ctx *gin.Context) {
	var req getBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	blog, err := server.store.GetBlog(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
