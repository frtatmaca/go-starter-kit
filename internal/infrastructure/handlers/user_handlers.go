package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-starter-kit/internal/application/ports"
)

type UserHandlers struct {
	userService ports.IUserService
}

var _ ports.IUserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(service ports.IUserService) *UserHandlers {
	return &UserHandlers{
		userService: service,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string

	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}

	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var confirmPassword string

	err := h.userService.Register(email, password, confirmPassword)
	if err != nil {
		return err
	}

	return nil
}
