package routers

import (
	controller_auth "github.com/servatom/diagon-alley/src/delivery/http/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func SetAuthRoutes(
	router *fiber.App,
	controller controller_auth.AuthController,
) {
	authRouter := router.Group("/auth")
	authRouter.Post("/login", controller.Login)
}