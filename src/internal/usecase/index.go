package usecase

import (
	domain_admin "github.com/servatom/diagon-alley/src/internal/domain/admin"
	domain_auth "github.com/servatom/diagon-alley/src/internal/domain/auth"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	domain_order "github.com/servatom/diagon-alley/src/internal/domain/order"
	"github.com/servatom/diagon-alley/src/internal/repository"
	usecase_admin "github.com/servatom/diagon-alley/src/internal/usecase/admin"
	usecase_auth "github.com/servatom/diagon-alley/src/internal/usecase/auth"
	usecase_product "github.com/servatom/diagon-alley/src/internal/usecase/product"
	usecase_order "github.com/servatom/diagon-alley/src/internal/usecase/order"
	"github.com/servatom/diagon-alley/src/utils"
)

type Usecases struct {
	Auth domain_auth.Usecase
	Product domain_product.Usecase
	Admin domain_admin.Usecase
	Order domain_order.Usecase
}

func InitUsecases(
	config *utils.Config,
	repositories *repository.Repositories,
) *Usecases {
	auth_usecase := usecase_auth.NewAuthUsecaseImplementation(config, repositories.Auth)
	product_usecase := usecase_product.NewProductUsecaseImplementation(config, repositories.Product)
	admin_usecase := usecase_admin.NewAdminUsecaseImplementation(config, product_usecase)
	order_usecase := usecase_order.NewOrderUsecaseImplementation(config, repositories.Order, repositories.OrderProductMap, product_usecase)
	return &Usecases{
		Auth: auth_usecase,
		Product: product_usecase,
		Admin: admin_usecase,
		Order: order_usecase,
	}
}