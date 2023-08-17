package controller_order

import (
	fiber "github.com/gofiber/fiber/v2"
	domain_order "github.com/servatom/diagon-alley/src/internal/domain/order"
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	utils "github.com/servatom/diagon-alley/src/utils"
)
type OrderController interface {
	CreateOrder (ctx *fiber.Ctx) (err error)
	GetAllOrders (ctx *fiber.Ctx) (err error)
}

type OrderControllerImplementation struct {
	config *utils.Config
	usecaseOrder domain_order.Usecase
}


func (o *OrderControllerImplementation) CreateOrder(
	ctx *fiber.Ctx,
) (err error) {
	var payload CreateOrderPayload
	err = ctx.BodyParser(&payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user := ctx.Locals("user").(*domain_auth.UserWithID)
	orderDetails, err := o.usecaseOrder.CreateOrder(ctx.Context(), user.ID, payload.ProductIDs)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(orderDetails)
}

func (o *OrderControllerImplementation) GetAllOrders(
	ctx *fiber.Ctx,
) (err error) {
	user := ctx.Locals("user").(*domain_auth.UserWithID)
	orders, err := o.usecaseOrder.GetAllOrders(ctx.Context(), user.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(orders)
}

func NewOrderControllerImplementation(
	config *utils.Config,
	usecaseOrder domain_order.Usecase,
) *OrderControllerImplementation {
	return &OrderControllerImplementation{
		config:      config,
		usecaseOrder: usecaseOrder,
	}
}