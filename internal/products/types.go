package products

import "context"

type ProductService interface {
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context) ([]*Product, error)
	CreateProduct(ctx context.Context, product *Product) error
}

type ProductStorage interface {
	Save(ctx context.Context, product *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context) ([]*Product, error)
}
