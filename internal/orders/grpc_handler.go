package orders

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "zenfulcode.com/commercifyms/api/order"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	svc OrderService
}

func NewHandler(server *grpc.Server, svc OrderService) {
	pb.RegisterOrderServiceServer(server, &grpcHandler{svc: svc})
}

func (h *grpcHandler) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func (h *grpcHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
