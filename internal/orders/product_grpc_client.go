package orders

import (
	"context"

	pb "github.com/zenfulcode/commercifyms/api/product"
)

type ProductGRPCClient struct {
	client pb.ProductServiceClient
}

func NewProductGRPCClient(client pb.ProductServiceClient) *ProductGRPCClient {
	return &ProductGRPCClient{
		client: client,
	}
}

func (c *ProductGRPCClient) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return c.client.ListProducts(ctx, req)
}
