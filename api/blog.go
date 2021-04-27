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
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// createBlog godoc
// @Summary Create a blog
// @Description Create blog by title and content
// @Tags blog
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param titleContent body createBlogRequest true "Blog request"
// @Success 200 {object} db.CreateBlogRow
// @Router /api/blogs/{id} [post]
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

// deleteBlog godoc
// @Summary Delete a blog
// @Description Delete blog by id
// @Tags blog
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {string} string "Blog deleted"
// @Router /api/blogs/{id} [delete]
func (server *Server) deleteBlog(ctx *gin.Context) {
	var req deleteBlogId
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
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Unauthorize Acces. Can't delete blog")))
		return
	}

	err = server.store.DeleteBlog(ctx, req.ID)
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

// updateBlog godoc
// @Summary Update a blog
// @Description Update blog title and content
// @Tags blog
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param titleContent body createBlogRequest true "Blog request"
// @Success 200 {string} string "Blog Updated"
// @Router /api/blogs/{id} [patch]
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

// getBlog godoc
// @Summary Show a blog
// @Description get blog by id
// @Tags blog
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} db.GetBlogRow
// @Router /api/blogs/{id} [get]
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

// listBlog godoc
// @Summary Show all blog
// @Description get all blog in json
// @Tags blog
// @Produce  json
// @Success 200 {object} []db.ListBlogRow
// @Router /api/blogs [get]
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
