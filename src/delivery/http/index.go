package delivery_rest

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	controller_admin "github.com/servatom/diagon-alley/src/delivery/http/controllers/admin"
	controller_auth "github.com/servatom/diagon-alley/src/delivery/http/controllers/auth"
	controller_order "github.com/servatom/diagon-alley/src/delivery/http/controllers/order"
	controller_product "github.com/servatom/diagon-alley/src/delivery/http/controllers/product"
	middleware_auth "github.com/servatom/diagon-alley/src/delivery/http/middleware"
	"github.com/servatom/diagon-alley/src/delivery/http/routers"
	"github.com/servatom/diagon-alley/src/internal/usecase"
	"github.com/servatom/diagon-alley/src/utils"
)

func RestDeliver(
	config *utils.Config,
	usecases *usecase.Usecases,
) {
	app := fiber.New()
	app.Use(cors.New())
	
	authController := controller_auth.NewAuthControllerImplementation(config, usecases.Auth)
	adminController := controller_admin.NewAdminControllerImplementation(config, usecases.Admin)
	productController := controller_product.NewProductControllerImplementation(config, usecases.Product)
	orderController := controller_order.NewOrderControllerImplementation(config, usecases.Order)
	
	auth_middleware := middleware_auth.NewAuthMiddlewareImplementation(config, usecases.Auth)
	routers.SetAuthRoutes(app, authController)
	routers.SetAdminRoutes(app, adminController, auth_middleware)
	routers.SetProductRoutes(app, productController)
	routers.SetOrderRoutes(app, orderController, auth_middleware)

	err := app.Listen(":9000")
	if err != nil {
		panic(err)
	}
}