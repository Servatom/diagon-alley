package repository_product

import (
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)
func (productRepo ProductRepository) toDomainProduct() *domain_product.ProductWithID{
	return &domain_product.ProductWithID{
		ID:       productRepo.ID,
		Product:     domain_product.Product{
			Name:    productRepo.Name,
			Price: productRepo.Price,
			Color: productRepo.Color,
			Image: productRepo.Image,
			AverageRating: productRepo.AverageRating,
			Description: productRepo.Description,
		},
	}
}
func NewProductRepository(
	product *domain_product.Product,
) *ProductRepository {
	return &ProductRepository{
		Name:        product.Name,
		Price:     product.Price,
		Color:     product.Color,
		Image:     product.Image,
		AverageRating:     product.AverageRating,
		Description:     product.Description,
	}
}
