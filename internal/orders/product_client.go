package orders

import (
	"context"

	pb "github.com/zenfulcode/commercifyms/api/product"
)

type ProductClient interface {
	ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error)
}
