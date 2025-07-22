package products

import (
	"context"

	pb "github.com/zenfulcode/commercifyms/api/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	pb.UnimplementedProductServiceServer
	svc ProductService
}

func NewGRPCHandler(server *grpc.Server, svc ProductService) {
	pb.RegisterProductServiceServer(server, &grpcHandler{
		svc: svc,
	})
}

func (h *grpcHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := h.svc.ListProducts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	var protoProducts []*pb.Product
	for _, p := range products {
		protoProducts = append(protoProducts, p.convertToProto())
	}

	return &pb.ListProductsResponse{Products: protoProducts}, nil
}
