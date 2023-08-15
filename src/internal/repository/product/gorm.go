package repository_product

import (
	"context"

	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	base_repository "github.com/servatom/diagon-alley/src/internal/repository/base"
	"github.com/servatom/diagon-alley/src/utils"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Name          string  `json:"name" gorm:"type:varchar(100);not null"`
	Price         float32 `json:"price" gorm:"type:float;not null"`
	Color         string  `json:"color" gorm:"type:varchar(100);not null"`
	Image         string  `json:"image" gorm:"type:varchar(100);not null"`
	AverageRating float32 `json:"average_rating" gorm:"type:float;not null"`
    Description   string  `json:"description" gorm:"type:varchar(100);not null"`
	base_repository.BaseRepository
}

type ProductRepositoryImplementation struct {
	db     *gorm.DB
	config *utils.Config
}

func (ProductRepository) TableName() string {
	return "product"
}

func (p *ProductRepositoryImplementation) CreateProduct (
	ctx context.Context,
	newProduct domain_product.Product,
) (*domain_product.ProductWithID, error) {
	newProductModel := NewProductRepository(&newProduct)
	err := p.db.Create(&newProductModel).Error
	if err != nil {
		return nil, err
	}
	return newProductModel.toDomainProduct(), nil
}

func (p *ProductRepositoryImplementation) UpdateProduct(
	ctx context.Context,
	newProduct domain_product.Product,
	productId int64,
) (*domain_product.ProductWithID, error) {
	newProductModel := NewProductRepository(&newProduct)
	err := p.db.Model(&ProductRepository{}).Where("id = ?", productId).Updates(&newProductModel).Error
	if err != nil {
		return nil, utils.HandleError(utils.ProductNotFound)
	}
	return newProductModel.toDomainProduct(), nil
}

func NewProductRepositoryImplementation(
	db *gorm.DB,
	config *utils.Config,
) *ProductRepositoryImplementation {
	err := db.AutoMigrate(&ProductRepository{})
	if err != nil {
		panic(err)
	}
	return &ProductRepositoryImplementation{
		db:     db,
		config: config,
	}
}
