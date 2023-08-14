package delivery_rest

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	controller_auth "github.com/servatom/diagon-alley/src/delivery/http/controllers/auth"
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
	routers.SetAuthRoutes(app, authController)
	err := app.Listen(":9000")
	if err != nil {
		panic(err)
	}
}