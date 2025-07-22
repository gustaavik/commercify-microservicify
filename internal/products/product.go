package products

import (
	pb "zenfulcode.com/commercifyms/api/product"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       int64
	Stock       int32
}

func (p *Product) convertToProto() *pb.Product {
	return &pb.Product{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
	}
}
