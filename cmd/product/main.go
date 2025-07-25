package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"zenfulcode.com/commercifyms/internal/products"
	"zenfulcode.com/commercifyms/pkg/common"
	"zenfulcode.com/commercifyms/pkg/db"
	"zenfulcode.com/commercifyms/pkg/trpc"

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
	if err := db.AutoMigrate(&products.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// --- 2. Connect to InventoryService via Consul ---
	// inventoryClient := trpc.ConnectToInventoryService()

	// --- 3. Setup gRPC server ---
	grpcServer := grpc.NewServer()

	// --- 4. Inject dependencies and register service ---
	productRepo := products.NewProductRepository(db)
	productService := products.NewProductService(&productRepo)

	products.NewGRPCHandler(grpcServer, productService)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	// --- 5. Start listener ---
	address := common.GetEnv("SERVICE_URL", ":60002")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// --- 6. Register with Consul ---
	trpc.RegisterWithConsul("product-service", address)

	log.Printf("ProductService running at %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
