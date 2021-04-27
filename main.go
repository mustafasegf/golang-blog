package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mustafasegf/golang-blog/api"
	db "github.com/mustafasegf/golang-blog/db/sqlc"
	"github.com/mustafasegf/golang-blog/util"
)

// @title Golang-blog project
// @version 1.0
// @description Simple golang blog for backend practice. This project created using gin and sqlc to generate type save code from sql queries.
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	server.SwaggerRouter()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}