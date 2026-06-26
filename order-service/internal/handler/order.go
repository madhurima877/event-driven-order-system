package handler

import (
	"encoding/json"
	"net/http"

	"github.com/madhurima877/order-service/internal/model"
	"github.com/madhurima877/order-service/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&order); err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
