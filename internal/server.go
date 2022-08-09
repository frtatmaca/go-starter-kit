package server

import (
	"github.com/gofiber/fiber/v2"
	"go-starter-kit/internal/application/ports"
	"go-starter-kit/internal/infrastructure/configuration"
	"log"
)

type Server struct {
	userHandlers         ports.IUserHandlers
	configurationManager configuration.IConfigurationManager
}

func NewServer(handlers ports.IUserHandlers, configurationManager configuration.IConfigurationManager) *Server {
	return &Server{
		userHandlers:         handlers,
		configurationManager: configurationManager,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)

	err := app.Listen(s.configurationManager.GetServerConfig().Port)
	if err != nil {
		log.Fatal(err)
	}
}
