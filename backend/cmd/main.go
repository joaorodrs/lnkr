package main

import (
	"log"
	"os"

	"github.com/joaorodrs/linker/internals/core/services"
	"github.com/joaorodrs/linker/internals/handlers"
	"github.com/joaorodrs/linker/internals/repositories"
	"github.com/joaorodrs/linker/internals/server"
)

func main() {
	mongoConn := os.Getenv("MONGO_URI")
	if mongoConn == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	//repositories
	linkRepository, err := repositories.NewLinkRepository(mongoConn)
	if err != nil {
		panic(err)
	}
	//services
	linkService := services.NewLinkService(linkRepository)
	//handlers
	linkHandlers := handlers.NewLinkHandlers(linkService)
	//server
	httpServer := server.NewServer(
		linkHandlers,
	)
	httpServer.Initialize()
}
