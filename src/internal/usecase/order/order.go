package usecase_order

import (
	"context"

	domain_order "github.com/servatom/diagon-alley/src/internal/domain/order"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	"github.com/servatom/diagon-alley/src/utils"
)

type OrderUsecaseImplementation struct {
	config *utils.Config
	orderRepo domain_order.OrderRepository
	orderProductRepo domain_order.OrderProductRepository
	productUsecase domain_product.Usecase
}

func (o *OrderUsecaseImplementation) CreateOrder(
	ctx context.Context,
	user_id int64,
	product_ids []int64,
) (*domain_order.OrderDetails, error) {
	order, err := o.orderRepo.CreateOrder(ctx, user_id)
	if err != nil {
		return nil, err
	}
	products, err := o.productUsecase.GetProductsByIds(ctx, product_ids)
	if err != nil {
		return nil, err
	}
	_, err = o.orderProductRepo.CreateOrderProductMap(ctx, order.ID, products)
	if err != nil {
		return nil, err
	}
	finalProducts := []domain_product.ProductWithID{}
	for _, product := range products {
		finalProducts = append(finalProducts, *product)
	}

	orderDetails := &domain_order.OrderDetails{
		ID: order.ID,
		Products: finalProducts,
	}

	return orderDetails, nil
}

func (o *OrderUsecaseImplementation) GetAllOrders(
	ctx context.Context,
	user_id int64,
) ([]*domain_order.OrderDetails, error) {
	orders, err := o.orderRepo.GetAllOrders(ctx, user_id)
	if err != nil {
		return nil, err
	}
	var finalOrders []*domain_order.OrderDetails
	for _, order := range orders {
		orderProductMaps, err := o.orderProductRepo.GetOrderProductMapByOrderId(ctx, order.ID)
		if err != nil {
			return nil, err
		}
		productIds := []int64{}
		for _, orderProductMap := range orderProductMaps {
			productIds = append(productIds, orderProductMap.ProductID)
		}
		products, err := o.productUsecase.GetProductsByIds(ctx, productIds)
		if err != nil {
			return nil, err
		}
		finalProducts := []domain_product.ProductWithID{}
		for _, product := range products {
			finalProducts = append(finalProducts, *product)
		}

		orderDetails := &domain_order.OrderDetails{
			ID: order.ID,
			Products: finalProducts,
		}
		finalOrders = append(finalOrders, orderDetails)
	}
	return finalOrders, nil
}

func NewOrderUsecaseImplementation(
	config *utils.Config,
	orderRepo domain_order.OrderRepository,
	orderProductRepo domain_order.OrderProductRepository,
	productUsecase domain_product.Usecase,
) *OrderUsecaseImplementation {
	return &OrderUsecaseImplementation{
		config: config,
		orderRepo: orderRepo,
		orderProductRepo: orderProductRepo,
		productUsecase: productUsecase,
	}
}