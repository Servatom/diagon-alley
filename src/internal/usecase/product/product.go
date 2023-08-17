package usecase_product

import (
	"context"

	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	"github.com/servatom/diagon-alley/src/utils"
)

type ProductUsecaseImplementation struct {
	config *utils.Config
	productRepo domain_product.Repository
}

func (p *ProductUsecaseImplementation) CreateProduct(
	ctx context.Context,
	newProduct domain_product.Product,
) (*domain_product.ProductWithID, error) {
	return p.productRepo.CreateProduct(ctx, newProduct)
}

func (p *ProductUsecaseImplementation) UpdateProduct(
	ctx context.Context,
	newProduct domain_product.Product,
	productId int64,
) (*domain_product.ProductWithID, error) {
	product, err := p.productRepo.UpdateProduct(ctx, newProduct, productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUsecaseImplementation) GetAllProducts(
	ctx context.Context,
) ([]*domain_product.ProductWithID, error) {
	products, err := p.productRepo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUsecaseImplementation) GetProductsByIds(
	ctx context.Context,
	product_ids []int64,
) ([]*domain_product.ProductWithID, error) {
	products, err := p.productRepo.GetProductsByIds(ctx, product_ids)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewProductUsecaseImplementation(
	config *utils.Config,
	productRepo domain_product.Repository,
) *ProductUsecaseImplementation {
	return &ProductUsecaseImplementation{
		config: config,
		productRepo: productRepo,
	}
}