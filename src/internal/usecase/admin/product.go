package usecase_admin

import (
	"context"

	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)

func (adminUsecase *AdminUsecaseImplementation) AddProduct(
	ctx context.Context,
	newProduct domain_product.Product,
) (*domain_product.ProductWithID, error) {
	product, err := adminUsecase.productUsecase.CreateProduct(ctx, newProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (adminUsecase *AdminUsecaseImplementation) UpdateProduct(
	ctx context.Context,
	product domain_product.Product,
	productId int64,
) (*domain_product.ProductWithID, error) {
	productWithID, err := adminUsecase.productUsecase.UpdateProduct(ctx, product, productId)
	if err != nil {
		return nil, err
	}
	return productWithID, nil
}