package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
	"github.com/mustafasegf/golang-blog/token"
	"github.com/mustafasegf/golang-blog/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	router     *gin.Engine
	db					db.Querier
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, db db.Querier) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		db: db,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)
	// router.GET("/users/:id", server.getAccount)

	// router.POST("/blogs", server.createBlog)
	// router.GET("/blogs/:id", server.getBlog)

	// router.POST("/blogs/:id/comments", server.createComment)
	// router.GET("/blogs/:id/comments", server.getBlog)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}