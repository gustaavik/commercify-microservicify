package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/zenfulcode/commercifyms/internal/orders"
	"github.com/zenfulcode/commercifyms/pkg/common"
	"github.com/zenfulcode/commercifyms/pkg/db"
	"github.com/zenfulcode/commercifyms/pkg/trpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// --- 1. Setup DB ---
	db := db.Init()

	if db == nil {
		panic("Failed to connect to the database")
	}

	// Ensure the database is migrated
	if err := db.AutoMigrate(&orders.Order{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// --- 2. Connect to ProductService via Consul ---
	productServiceClient := trpc.ConnectToProductService()
	productClient := orders.NewProductGRPCClient(productServiceClient)

	// --- 3. Setup gRPC server ---
	grpcServer := grpc.NewServer()

	// --- 4. Inject dependencies and register service ---
	orderRepo := orders.NewGormOrderRepository(db)
	orderService := orders.NewOrderService(productClient, orderRepo)

	orders.NewHandler(grpcServer, orderService)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// --- 5. Start listener ---
	address := common.GetEnv("SERVICE_URL", ":60001")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// --- 6. Register with Consul ---
	trpc.RegisterWithConsul("order-service", address)

	log.Printf("OrderService running at %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
