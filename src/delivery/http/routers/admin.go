package routers

import (
	"github.com/gofiber/fiber/v2"
	controller_admin "github.com/servatom/diagon-alley/src/delivery/http/controllers/admin"
	middleware_auth "github.com/servatom/diagon-alley/src/delivery/http/middleware"
)

func SetAdminRoutes(
	router *fiber.App,
	controller controller_admin.AdminController,
	authMiddleware *middleware_auth.AuthMiddlewareImplementation,
) {
	productAdminRouter := router.Group("/admin/product")
	productAdminRouter.Use(authMiddleware.VerifyTokenForAdmin)
	productAdminRouter.Post("/", controller.AddProduct)
	productAdminRouter.Patch("/update/:productID", controller.UpdateProduct)
}
