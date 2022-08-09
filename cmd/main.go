package main

import (
	server "go-starter-kit/internal"
	"go-starter-kit/internal/application/repository"
	"go-starter-kit/internal/infrastructure/configuration"
	"go-starter-kit/internal/infrastructure/handlers"
	"go-starter-kit/internal/infrastructure/services"
)

func main() {
	config := configuration.NewConfigurationManager("./internal/infrastructure/resource", "application")

	userRepository := repository.NewUserRepository(config.GetMongoDbConfig().ConnectionString)
	userService := services.NewUserService(userRepository)
	userHandlers := handlers.NewUserHandlers(userService)

	httpServer := server.NewServer(userHandlers, config)
	httpServer.Initialize()
}
