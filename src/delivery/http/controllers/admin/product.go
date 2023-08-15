package controller_admin

import (
	"github.com/gofiber/fiber/v2"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)

func (controller *AdminControllerImplementation) AddProduct(
	ctx *fiber.Ctx,
) (err error){
	payload := new(ProductPayload)
	if err := ctx.BodyParser(payload); err != nil {
		return err
	}
	product, err := controller.usecaseAdmin.AddProduct(ctx.Context(), (domain_product.Product)(*payload))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": product,
	})
}

func (controller *AdminControllerImplementation) UpdateProduct(
	ctx *fiber.Ctx,
) (err error){
	productID, err := ctx.ParamsInt("productID")
	if err != nil {
		return err
	}
	payload := new(ProductPayload)
	if err := ctx.BodyParser(payload); err != nil {
		return err
	}
	product, err := controller.usecaseAdmin.UpdateProduct(ctx.Context(), (domain_product.Product)(*payload), (int64)(productID))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"product": product,
	})
}


