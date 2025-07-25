package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"zenfulcode.com/commercifyms/internal/gateway"
	"zenfulcode.com/commercifyms/pkg/common"
	"zenfulcode.com/commercifyms/pkg/trpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to internal services
	productClient := trpc.ConnectToProductService()
	orderClient := trpc.ConnectToOrderService()

	mux := http.NewServeMux()

	handler := gateway.NewHandler(orderClient, productClient)
	handler.RegisterRoutes(mux)

	httpAddress := common.GetEnv("GATEWAY_SERVICE_URL", ":6091")

	log.Printf("Gateway listening on %s", httpAddress)

	if err := http.ListenAndServe(httpAddress, mux); err != nil {
		panic(err)
	}
}
