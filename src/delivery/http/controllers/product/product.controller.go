package controller_product

import (
	fiber "github.com/gofiber/fiber/v2"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	utils "github.com/servatom/diagon-alley/src/utils"
)
type ProductController interface {
	CreateProduct (ctx *fiber.Ctx) (err error)
	UpdateProduct (ctx *fiber.Ctx) (err error)
}

type ProductControllerImplementation struct {
	config *utils.Config
	usecaseProduct domain_product.Usecase
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