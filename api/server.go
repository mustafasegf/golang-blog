package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
	"github.com/mustafasegf/golang-blog/token"
	"github.com/mustafasegf/golang-blog/util"	

	"github.com/mustafasegf/golang-blog/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	apiAuth := router.Group("/api").Use(authMiddleware(server.tokenMaker))

	api.POST("/users", server.createUser)
	api.POST("/users/login", server.loginUser)

	apiAuth.POST("/blogs", server.createBlog)
	api.GET("/blogs", server.listBlog)
	api.GET("/blogs/:id", server.getBlog)

	apiAuth.PATCH("/blogs/:id", server.updateBlog)
	apiAuth.DELETE("/blogs/:id", server.deleteBlog)

	apiAuth.POST("/blogs/:id/comments", server.createComment)
	api.GET("/blogs/:id/comments", server.getComment)

	apiAuth.PATCH("/comments/:id", server.updateComment)
	apiAuth.DELETE("/comments/:id", server.deleteComment)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) SwaggerRouter() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "API Payment"
	docs.SwaggerInfo.Description = "Payment Application Programming Interface"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = server.config.ServerAddress
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
