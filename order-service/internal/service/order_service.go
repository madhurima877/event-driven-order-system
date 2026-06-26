package service

import (
	"github.com/madhurima877/order-service/internal/model"
	"github.com/madhurima877/order-service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService)Create(order *model.Order)error{
	return s.repo.Create(order)
}