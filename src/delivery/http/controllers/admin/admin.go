package controller_admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/servatom/diagon-alley/src/utils"
	domain_admin "github.com/servatom/diagon-alley/src/internal/domain/admin"
)

type AdminController interface {
	AddProduct (ctx *fiber.Ctx) (err error)
	UpdateProduct (ctx *fiber.Ctx) (err error)
}

type AdminControllerImplementation struct {
	config       *utils.Config
	usecaseAdmin domain_admin.Usecase
}

func NewAdminControllerImplementation(
	config *utils.Config,
	usecaseAdmin domain_admin.Usecase,
) *AdminControllerImplementation {
	return &AdminControllerImplementation{
		config:       config,
		usecaseAdmin: usecaseAdmin,
	}
}
