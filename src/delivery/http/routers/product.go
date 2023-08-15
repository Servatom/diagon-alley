package routers

import (
	controller_product "github.com/servatom/diagon-alley/src/delivery/http/controllers/product"
	"github.com/gofiber/fiber/v2"
)

func SetProductRoutes(
	router *fiber.App,
	controller controller_product.ProductController,
) {
}
