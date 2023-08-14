package middleware_auth

import (
	"strings"

	"errors"

	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	utils "github.com/servatom/diagon-alley/src/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleWare interface {
	VerifyTokenForUser(ctx *fiber.Ctx) (err error)
	VerifyTokenForAdmin(ctx *fiber.Ctx) (err error)
}

type AuthMiddlewareImplementation struct {
	config      *utils.Config
	authUsecase domain_auth.Usecase
}

func (am *AuthMiddlewareImplementation) VerifyToken(
	ctx *fiber.Ctx,
) (*domain_auth.UserWithID, error) {
	tokenString := ctx.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	} else {
		return nil, errors.New("Unauthorized")
	}
	if tokenString == "" {
		return nil, errors.New("Unauthorized")
	}
	user, err := am.authUsecase.VerifyToken(ctx.Context(), tokenString)
	if err != nil {
		return nil, errors.New("Unauthorized")
	}
	return user, nil
}

func (am *AuthMiddlewareImplementation) VerifyTokenForUser(
	ctx *fiber.Ctx,
) error {
	user, err := am.VerifyToken(ctx)
	if err != nil {
		return ctx.Redirect("/auth/login")
	}
	ctx.Locals("user", user)
	return ctx.Next()
}

func (am *AuthMiddlewareImplementation) VerifyTokenForAdmin(
	ctx *fiber.Ctx,
) (err error) {
	user, err := am.VerifyToken(ctx)
	if err != nil {
		return ctx.Redirect("/auth/login")
	}
	if !user.IsAdmin {
		ctx.Status(403).SendString("Unauthorised")
		return errors.New("Unauthorised")
	}
	ctx.Locals("user", user)
	return ctx.Next()
}

func NewAuthMiddlewareImplementation(
	config *utils.Config,
	authUsecase domain_auth.Usecase,
) *AuthMiddlewareImplementation {
	return &AuthMiddlewareImplementation{
		config:      config,
		authUsecase: authUsecase,
	}
}
