package routers

import (
	"github.com/gofiber/fiber/v2"
)

func SetHealthRoutes(
	router *fiber.App,
) {
	healthRouter := router.Group("/health")
	healthRouter.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server is up and running!",
		})
	})
}
