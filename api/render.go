package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (server *Server) index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}

func (server *Server) renderBlog(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "blog.html", gin.H{"id": ctx.Param("id")})
}

func (server *Server) register(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", "")
}

func (server *Server) login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", "")
}