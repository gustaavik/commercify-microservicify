package products

import (
	"context"

	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) store {
	return store{db: db}
}

func (s *store) Save(ctx context.Context, product *Product) error {
	return s.db.WithContext(ctx).Create(&product).Error
}

func (s *store) GetByID(ctx context.Context, id string) (*Product, error) {
	var product Product
	if err := s.db.WithContext(ctx).First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *store) List(ctx context.Context) ([]*Product, error) {
	var products []*Product
	if err := s.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
