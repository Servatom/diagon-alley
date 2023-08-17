package repository_order

import (
	domain_order "github.com/servatom/diagon-alley/src/internal/domain/order"
)

func(orderRepo OrderRepository) toDomainOrder() *domain_order.Order{
	return &domain_order.Order{
		ID: orderRepo.ID,
		UserID: orderRepo.UserID,
	}
}

func (orderProductRepo OrderProductMapRepository) toDomainOrderProductMap() *domain_order.OrderProductMap{
	return &domain_order.OrderProductMap{
		OrderID: orderProductRepo.OrderID,
		ProductID: orderProductRepo.ProductID,
	}
}