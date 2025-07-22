BINS := product order gateway
BIN_DIR := ./bin

build:
	@mkdir -p $(BIN_DIR)
	@for svc in $(BINS); do \
		go build -o $(BIN_DIR)/$$svc ./cmd/$$svc; \
	done

run-grid:
	@tmux new-session -d -s commercify-dev -n services

	# Pane 0: Start Consul agent in dev mode
	@tmux send-keys -t commercify-dev:0 'consul agent -dev -client=0.0.0.0' C-m

	# Wait until Consul is up (check HTTP API port 8500)
	@sleep 2
	@echo "Waiting for Consul to become available..."
	@until curl -s http://localhost:8500/v1/status/leader | grep -q '"'; do \
		sleep 1; \
	done
	@echo "Consul is up!"

	# Pane 1: Product service
	@tmux split-window -h -t commercify-dev:0
	@tmux split-window -v -t commercify-dev:0.0
	@tmux split-window -v -t commercify-dev:0.1
	@tmux send-keys -t commercify-dev:0.1 'cd cmd/product && air' C-m

	# Pane 2: Order service
	@tmux send-keys -t commercify-dev:0.2 'cd cmd/order && air' C-m

	# Pane 3: Gateway service
	@tmux send-keys -t commercify-dev:0.3 'cd cmd/gateway && air' C-m

	# Layout and attach
	@tmux select-layout -t commercify-dev:0 tiled
	@tmux select-pane -t commercify-dev:0
	@tmux attach -t commercify-dev

build-product:
	go build -o bin/product ./cmd/product

build-order:
	go build -o bin/order ./cmd/order

build-gateway:
	go build -o bin/gateway ./cmd/gateway

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

run-order:
	go run cmd/order/main.go

run-product:
	go run cmd/product/main.go

proto:
	@protoc -I=proto \
		--go_out=api --go-grpc_out=api \
		proto/*.proto