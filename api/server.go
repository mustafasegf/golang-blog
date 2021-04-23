package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
	"github.com/mustafasegf/golang-blog/token"
	"github.com/mustafasegf/golang-blog/util"
)

// Server serves HTTP requests for blog service.
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
	store      db.Store
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		store:      store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	
	router.GET("/", server.index)
	router.GET("/register", server.register)
	router.GET("/login", server.login)
	router.GET("blogs/:id", server.renderBlog)

	api := router.Group("/api")


	api.POST("/users", server.createUser)
	api.POST("/users/login", server.loginUser)

	apiAuth := api.Use(authMiddleware(server.tokenMaker))

	apiAuth.POST("/blogs", server.createBlog)
	apiAuth.GET("/blogs", server.listBlog)
	apiAuth.GET("/blogs/:id", server.getBlog)

	apiAuth.POST("/blogs/:id/update", server.updateBlog)
	apiAuth.POST("/blogs/:id/delete", server.deleteBlog)

	apiAuth.POST("/blogs/:id/comments", server.createComment)
	apiAuth.GET("/blogs/:id/comments", server.getComment)

	apiAuth.POST("/comments/:id/update", server.updateComment)
	apiAuth.POST("/comments/:id/delete", server.deleteComment)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
