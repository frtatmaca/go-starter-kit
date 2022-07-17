package server

import (
	"github.com/gofiber/fiber/v2"
	"go-starter-kit/internal/application/ports"
	"log"
)

type Server struct {
	userHandlers ports.IUserHandlers
}

func NewServer(handlers ports.IUserHandlers) *Server {
	return &Server{
		userHandlers: handlers,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
