package main

import (
	server "go-starter-kit/internal"
	"go-starter-kit/internal/application/repository"
	"go-starter-kit/internal/infrastructure/handlers"
	"go-starter-kit/internal/infrastructure/services"
)

func main() {
	mongoConn := "connection string"
	userRepository := repository.NewUserRepository(mongoConn)
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandlers(userService)

	httpServer := server.NewServer(userHandlers)
	httpServer.Initialize()
}
