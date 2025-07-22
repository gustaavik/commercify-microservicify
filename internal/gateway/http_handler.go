package gateway

import (
	"net/http"

	orderpb "zenfulcode.com/commercifyms/api/order"
	productpb "zenfulcode.com/commercifyms/api/product"
	"zenfulcode.com/commercifyms/pkg/common"
)

type handler struct {
	orderClient   orderpb.OrderServiceClient
	productClient productpb.ProductServiceClient
}

func NewHandler(orderClient orderpb.OrderServiceClient, productClient productpb.ProductServiceClient) *handler {
	return &handler{
		orderClient,
		productClient,
	}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/orders", h.handleCreateOrder)
}

func (h *handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orderpb.CreateOrderRequest
	if err := common.ReadJSON(r, &req); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.orderClient.CreateOrder(r.Context(), &req)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusCreated, result)
}
