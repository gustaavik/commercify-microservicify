package orders

import (
	"context"

	pb "zenfulcode.com/commercifyms/api/product"
)

type ProductClient interface {
	ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error)
}
