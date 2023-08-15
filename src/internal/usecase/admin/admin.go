package usecase_admin

import (
	"github.com/servatom/diagon-alley/src/utils"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)

type AdminUsecaseImplementation struct {
	config *utils.Config
	productUsecase domain_product.Usecase
}

func NewAdminUsecaseImplementation(
	config *utils.Config,
	productUsecase domain_product.Usecase,
) *AdminUsecaseImplementation {
	return &AdminUsecaseImplementation{
		config: config,
		productUsecase: productUsecase,
	}
}