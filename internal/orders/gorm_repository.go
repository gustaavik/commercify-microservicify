package orders

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, order Order) error {
	return r.db.WithContext(ctx).Create(&order).Error
}

func (r *repository) Get(ctx context.Context, id string) (Order, error) {
	var order Order
	err := r.db.WithContext(ctx).First(&order, "id = ?", id).Error
	if err != nil {
		return Order{}, err
	}
	return order, nil
}

func (r *repository) List(ctx context.Context) ([]Order, error) {
	var orders []Order
	err := r.db.WithContext(ctx).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
