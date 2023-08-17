package routers

import (
	controller_order "github.com/servatom/diagon-alley/src/delivery/http/controllers/order"
	"github.com/gofiber/fiber/v2"
	middleware_auth "github.com/servatom/diagon-alley/src/delivery/http/middleware"
)

func SetOrderRoutes(
	router *fiber.App,
	controller controller_order.OrderController,
	authMiddleware *middleware_auth.AuthMiddlewareImplementation,
) {
	orderRouter := router.Group("/order")
	orderRouter.Use(authMiddleware.VerifyTokenForUser)
	orderRouter.Post("/create", controller.CreateOrder)
	orderRouter.Get("/all", controller.GetAllOrders)
}
