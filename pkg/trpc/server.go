package trpc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hashicorp/consul/api"
	orderpb "zenfulcode.com/commercifyms/api/order"
	productpb "zenfulcode.com/commercifyms/api/product"
	"zenfulcode.com/commercifyms/pkg/common"
)

func connectToService(serviceName string) *grpc.ClientConn {
	host := common.GetEnv("CONSUL_HTTP_ADDR", ":8500")

	conn, err := grpc.NewClient(
		fmt.Sprintf("consul://%s/%s?wait=14s", host, serviceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", serviceName, err)
	}
	return conn
}

func ConnectToProductService() productpb.ProductServiceClient {
	return productpb.NewProductServiceClient(connectToService("product-service"))
}

func ConnectToOrderService() orderpb.OrderServiceClient {
	return orderpb.NewOrderServiceClient(connectToService("order-service"))
}

func RegisterWithConsul(name, clientHost string) {
	config := api.DefaultConfig()
	config.Address = common.GetEnv("CONSUL_HTTP_ADDR", ":8500")

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("consul client error: %v", err)
	}

	address := strings.Split(clientHost, ":")[0]
	if address == "" {
		address = "localhost"
	}
	port, err := strconv.Atoi(strings.Split(clientHost, ":")[1])
	if err != nil {
		log.Fatalf("failed to parse port: %v", err)
	}

	registration := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%d", name, port),
		Name:    name,
		Address: address, // should match Docker service name or hostname
		Port:    port,
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", address, port),
			Interval:                       "10s",
			Timeout:                        "3s",
			DeregisterCriticalServiceAfter: "1m",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		log.Fatalf("failed to register service with consul: %v", err)
	}
}
