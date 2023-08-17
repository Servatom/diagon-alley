package domain_order

import (
	"context"

	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
)

type Order struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type OrderProductMap struct {
	OrderID int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
}

type OrderDetails struct {
	ID int64 `json:"id"`
	Products []domain_product.ProductWithID `json:"products"`
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, user_id int64) (*Order, error)
	GetAllOrders(ctx context.Context, user_id int64) ([]*Order, error)
}

type OrderProductRepository interface {
	CreateOrderProductMap(ctx context.Context, order_id int64, products []*domain_product.ProductWithID) (bool, error)
	GetOrderProductMapByOrderId(ctx context.Context, order_id int64) ([]*OrderProductMap, error)
}

type Usecase interface {
	CreateOrder(ctx context.Context, user_id int64, product_ids []int64) (*OrderDetails, error)
	GetAllOrders(ctx context.Context, user_id int64) ([]*OrderDetails, error)
}