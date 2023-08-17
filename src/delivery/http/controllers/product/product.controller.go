package controller_product

import (
	fiber "github.com/gofiber/fiber/v2"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	utils "github.com/servatom/diagon-alley/src/utils"
)
type ProductController interface {
	GetAllProducts (ctx *fiber.Ctx) (err error)
}

type ProductControllerImplementation struct {
	config *utils.Config
	usecaseProduct domain_product.Usecase
}

func (controller *ProductControllerImplementation) GetAllProducts(
	ctx *fiber.Ctx,
) (err error) {
	products, err := controller.usecaseProduct.GetAllProducts(ctx.Context())
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}


func NewProductControllerImplementation(
	config *utils.Config,
	usecaseProduct domain_product.Usecase,
) *ProductControllerImplementation {
	return &ProductControllerImplementation{
		config:      config,
		usecaseProduct: usecaseProduct,
	}
}