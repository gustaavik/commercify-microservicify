package products

import (
	"context"
)

type service struct {
	repo ProductStorage
}

func NewProductService(repo ProductStorage) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProduct(ctx context.Context, product *Product) error {
	if err := s.repo.Save(ctx, product); err != nil {
		return err
	}
	return nil
}

func (s *service) GetProductByID(ctx context.Context, id string) (*Product, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) ListProducts(ctx context.Context) ([]*Product, error) {
	products, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
