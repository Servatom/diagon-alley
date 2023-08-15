package domain_admin

import (
	"context"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)

type Usecase interface {
	AddProduct (ctx context.Context, newProduct domain_product.Product) (*domain_product.ProductWithID, error)
	UpdateProduct (ctx context.Context, product domain_product.Product, productId int64) (*domain_product.ProductWithID, error)
}