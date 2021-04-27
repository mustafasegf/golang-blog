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

type createCommentsRequest struct {
	BlogID  int32  `uri:"id"`
	Comment string `json:"comment" binding:"required"`
}

// createComment godoc
// @Summary Create a comment
// @Description Create comment by id and comment
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param content body string true "Comment content"
// @Success 200 {object} db.CreateCommentRow
// @Param Authorization header string true "Bearer Token"
// @Router /api/blogs/{id}/comments [post]
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateCommentParams{
		BlogID:  req.BlogID,
		UserID:  authPayload.UserId,
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

// getComment godoc
// @Summary get all comment in blog
// @Description get all comment by blog id
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} db.CreateCommentRow
// @Router /api/comments/{id} [get]
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

type updateCommentRequest struct {
	ID      int32  `uri:"id"`
	Comment string `json:"comment" binding:"required"`
}

// updateComment godoc
// @Summary update comment 
// @Description update comment by comment id
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param content body string true "Comment content"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} db.CreateCommentRow
// @Router /api/comments/{id} [patch]
func (server *Server) updateComment(ctx *gin.Context) {
	var req updateCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment, err := server.store.GetOneComment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.UserId != comment.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorize Acces. Can't change comment")))
		return
	}
	arg := db.UpdateCommentParams{
		ID:      req.ID,
		Comment: req.Comment,
	}
	err = server.store.UpdateComment(ctx, arg)
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

	ctx.JSON(http.StatusOK, "Comment Updated")
}

type deleteCommentId struct {
	ID int32 `uri:"id" binding:"required"`
}

// deleteComment godoc
// @Summary delete comment 
// @Description delete comment by comment id
// @Tags comment
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} db.CreateCommentRow
// @Router /api/comments/{id} [delete]
func (server *Server) deleteComment(ctx *gin.Context) {
	var req deleteCommentId
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comment, err := server.store.GetOneComment(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.UserId != comment.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorize Acces. Can't delete comment")))
		return
	}

	err = server.store.DeleteComment(ctx, req.ID)
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

	ctx.JSON(http.StatusOK, "Comment deleted")
}
