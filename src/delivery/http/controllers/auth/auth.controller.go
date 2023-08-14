package controller_auth

import (
	fiber "github.com/gofiber/fiber/v2"
	utils "github.com/servatom/diagon-alley/src/utils"
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
)
type AuthController interface {
	Login(ctx *fiber.Ctx) (err error)
}

type AuthControllerImplementation struct {
	config *utils.Config
	usecaseAuth domain_auth.Usecase
}

func (ac *AuthControllerImplementation) Login(ctx *fiber.Ctx) (err error) {
	var payload LoginPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return err
	}
	token, err := ac.usecaseAuth.Login(ctx.Context(), payload.Email, payload.Password)
	if err != nil {
		return err
	}
	return ctx.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

func NewAuthControllerImplementation(
	config *utils.Config,
	usecaseAuth domain_auth.Usecase,
) *AuthControllerImplementation {
	return &AuthControllerImplementation{
		config:      config,
		usecaseAuth: usecaseAuth,
	}
}