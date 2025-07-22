package orders

import "context"

type service struct {
	productClient ProductClient
	repo          OrderStorage
}

func NewOrderService(productClient ProductClient, repo OrderStorage) *service {
	return &service{
		productClient: productClient,
		repo:          repo,
	}
}

func (s *service) CreateOrder(ctx context.Context, order Order) error {
	// Here you would typically validate the order, check product availability, etc.
	// For simplicity, we just save the order.
	return s.repo.Save(ctx, order)
}
