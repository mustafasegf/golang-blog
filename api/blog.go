package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
	"github.com/mustafasegf/golang-blog/token"
)

type createBlogRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

func (server *Server) createBlog(ctx *gin.Context) {
	var req createBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateBlogParams{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: authPayload.UserId,
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

	ctx.JSON(http.StatusOK, blog)
}

type deleteBlogId struct {
	ID int32 `uri:"id" binding:"required"`
}

type deleteTokenId struct {
	Token string `json:"token" binding:"required"`
}

func (server *Server) deleteBlog(ctx *gin.Context) {
	var reqId deleteBlogId
	if err := ctx.ShouldBindUri(&reqId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var reqToken deleteTokenId
	if err := ctx.ShouldBindJSON(&reqToken); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	blog, err := server.store.GetBlog(ctx, reqId.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.UserId != blog.Userid {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorize Acces. Can't delete blog")))
		return
	}

	err = server.store.DeleteBlog(ctx, reqId.ID)
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

	ctx.JSON(http.StatusOK, "Blog deleted")
}

type updateBlogRequest struct {
	ID      int32  `uri:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (server *Server) updateBlog(ctx *gin.Context) {
	var req updateBlogRequest
	var err error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.UserId != blog.Userid {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorize Acces. Can't change blog")))
		return
	}

	if req.Content != "" && req.Title != "" {
		arg := db.UpdateBlogParams{
			ID:      req.ID,
			Title:   req.Title,
			Content: req.Content,
		}
		err = server.store.UpdateBlog(ctx, arg)
	} else if req.Title != "" {
		arg := db.UpdateBlogTitleParams{
			ID:    req.ID,
			Title: req.Title,
		}
		err = server.store.UpdateBlogTitle(ctx, arg)
	} else if req.Content != "" {
		arg := db.UpdateBlogContentParams{
			ID:      req.ID,
			Content: req.Content,
		}
		err = server.store.UpdateBlogContent(ctx, arg)
	} else {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
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

	ctx.JSON(http.StatusOK, "Blog Updated")
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
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (server *Server) listBlog(ctx *gin.Context) {
	blog, err := server.store.ListBlog(ctx)
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
