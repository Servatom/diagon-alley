package routers

import (
	"github.com/gofiber/fiber/v2"
	controller_auth "github.com/servatom/diagon-alley/src/delivery/http/controllers/auth"
	middleware_auth "github.com/servatom/diagon-alley/src/delivery/http/middleware"
)

func SetAuthRoutes(
	router *fiber.App,
	controller controller_auth.AuthController,
	authMiddleware *middleware_auth.AuthMiddlewareImplementation,
) {
	authRouter := router.Group("/auth")
	authRouter.Post("/login", controller.Login)

	userRouter := router.Group("/auth/user")
	userRouter.Use(authMiddleware.VerifyTokenForUser)
	userRouter.Get("/me", controller.Me)
}