package domain_product 

import (
	"context"

)

type Product struct {
	Name string `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
	Color string `json:"color,omitempty"`
	Image string `json:"image,omitempty"`
	AverageRating float32 `json:"average_rating,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProductWithID struct {
	ID int64 `json:"id"`
	Product `json:",inline"`    
}

type Repository interface {
	CreateProduct(ctx context.Context, newProduct Product) (*ProductWithID, error)
	UpdateProduct(ctx context.Context, newProduct Product, productId int64) (*ProductWithID, error)
	GetAllProducts(ctx context.Context) ([]*ProductWithID, error)
	GetProductsByIds(ctx context.Context, productIds []int64) ([]*ProductWithID, error)
}

type Usecase interface {
	CreateProduct(ctx context.Context, newProduct Product) (*ProductWithID, error)
	UpdateProduct(ctx context.Context, newProduct Product, productId int64) (*ProductWithID, error)
	GetAllProducts(ctx context.Context) ([]*ProductWithID, error)
	GetProductsByIds(ctx context.Context, productIds []int64) ([]*ProductWithID, error)
}