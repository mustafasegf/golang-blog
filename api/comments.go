package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
)

type createCommentsRequest struct {
	BlogID  int32  `uri:"id"`
	UserID  int32  `json:"userId" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	Token   string `json:"token" binding:"required"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if _, err := server.tokenMaker.VerifyToken(req.Token); err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		BlogID:  req.BlogID,
		UserID:  req.UserID,
		Comment: req.Comment,
	}

	comment, err := server.store.CreateComment(ctx, arg)
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

	ctx.JSON(http.StatusOK, comment)
}

type getCommentRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getComment(ctx *gin.Context) {
	var req getBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment, err := server.store.GetComment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, comment)
}
