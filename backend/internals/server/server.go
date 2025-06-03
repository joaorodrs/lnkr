package server

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joaorodrs/linker/internals/core/ports"
)

type Server struct {
	//We will add every new Handler here
	linkHandlers ports.LinkHandlers
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(lHandlers ports.LinkHandlers) *Server {
	return &Server{
		linkHandlers: lHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	linkRoutes := v1.Group("/link")
	linkRoutes.Post("/", s.linkHandlers.CreateLink)
	linkRoutes.Get("/", s.linkHandlers.GetAllLinks)
	linkRoutes.Get("/r/:hash", s.linkHandlers.GetLink)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
