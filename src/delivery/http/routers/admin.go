package routers

import (
	"github.com/gofiber/fiber/v2"
	controller_admin "github.com/servatom/diagon-alley/src/delivery/http/controllers/admin"
)

func SetAdminRoutes(
	router *fiber.App,
	controller controller_admin.AdminController,
) {
	productAdminRouter := router.Group("/admin/product")
	productAdminRouter.Post("/add", controller.AddProduct)
	productAdminRouter.Patch("/update/:productID", controller.UpdateProduct)
}
