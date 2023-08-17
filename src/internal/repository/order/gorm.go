package repository_order

import (
	"context"

	domain_order "github.com/servatom/diagon-alley/src/internal/domain/order"
	domain_product "github.com/servatom/diagon-alley/src/internal/domain/product"
	base_repository "github.com/servatom/diagon-alley/src/internal/repository/base"
	"github.com/servatom/diagon-alley/src/utils"
	"gorm.io/gorm"
)

type OrderRepository struct {
	UserID int64 `json:"user_id" gorm:"not null"`
	OrderProductMaps []OrderProductMapRepository `json:"order_product_maps" gorm:"foreignKey:OrderID"`
	base_repository.BaseRepository
}

type OrderProductMapRepository struct {
	OrderID int64 `json:"order_id" gorm:"not null"`
	ProductID int64 `json:"product_id" gorm:"not null"`
	base_repository.BaseRepository
}

type OrderRepositoryImplementation struct {
	db     *gorm.DB
	config *utils.Config
}

type OrderProductMapRepositoryImplementation struct {
	db     *gorm.DB
	config *utils.Config
}

func (OrderRepository) TableName() string {
	return "order"
}

func (OrderProductMapRepository) TableName() string {
	return "order_product_map"
}

func (o *OrderRepositoryImplementation) CreateOrder(
	ctx context.Context,
	user_id int64,
) (*domain_order.Order, error) {
	newOrder := &OrderRepository{
		UserID: user_id,
	}
	err := o.db.Create(&newOrder).Error
	if err != nil {
		return nil, err
	}
	return newOrder.toDomainOrder(), nil
}

func (o *OrderRepositoryImplementation) GetAllOrders(
	ctx context.Context,
	user_id int64,
) ([]*domain_order.Order, error) {
	orders := []*OrderRepository{}
	err := o.db.Where("user_id = ?", user_id).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var finalOrders []*domain_order.Order
	for _, order := range orders {
		finalOrders = append(finalOrders, order.toDomainOrder())
	}
	return finalOrders, nil
}

func (o *OrderProductMapRepositoryImplementation) CreateOrderProductMap(
	ctx context.Context,
	order_id int64,
	products []*domain_product.ProductWithID,
) (bool, error) {
	for _, product := range products {
		newOrderProductMap := &OrderProductMapRepository{
			OrderID: order_id,
			ProductID: product.ID,
		}
		err := o.db.Create(&newOrderProductMap).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (o *OrderProductMapRepositoryImplementation) GetOrderProductMapByOrderId(
	ctx context.Context,
	order_id int64,
) ([]*domain_order.OrderProductMap, error) {
	orderProductMaps := []*OrderProductMapRepository{}
	err := o.db.Where("order_id = ?", order_id).Find(&orderProductMaps).Error
	if err != nil {
		return nil, err
	}
	var finalOrderProductMaps []*domain_order.OrderProductMap
	for _, orderProductMap := range orderProductMaps {
		finalOrderProductMaps = append(finalOrderProductMaps, orderProductMap.toDomainOrderProductMap())
	}
	return finalOrderProductMaps, nil
}

func NewOrderRepositoryImplementation(
	db *gorm.DB,
	config *utils.Config,
) *OrderRepositoryImplementation {
	err := db.AutoMigrate(&OrderRepository{})
	if err != nil {
		panic(err)
	}
	return &OrderRepositoryImplementation{
		db:     db,
		config: config,
	}
}

func NewOrderProductMapRepositoryImplementation(
	db *gorm.DB,
	config *utils.Config,
) *OrderProductMapRepositoryImplementation {
	err := db.AutoMigrate(&OrderProductMapRepository{})
	if err != nil {
		panic(err)
	}
	return &OrderProductMapRepositoryImplementation{
		db:     db,
		config: config,
	}
}
