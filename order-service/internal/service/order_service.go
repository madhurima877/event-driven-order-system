package service

import (
	"encoding/json"
	"log"

	"github.com/madhurima877/order-service/internal/kafka"
	"github.com/madhurima877/order-service/internal/model"
	"github.com/madhurima877/order-service/internal/repository"

	kafka_go "github.com/segmentio/kafka-go"
)

type OrderService struct {
	repo   *repository.OrderRepository
	writer *kafka_go.Writer
}

func NewOrderService(repo *repository.OrderRepository, writer *kafka_go.Writer) *OrderService {
	return &OrderService{
		repo:   repo,
		writer: writer,
	}
}

func (s *OrderService) Create(order *model.Order) error {
	if err := s.repo.Create(order); err != nil {
		return err
	}
	message, _ := json.Marshal(order)
	err := kafka.Publish(s.writer, message)
	if err != nil {
		return err
	}
	log.Println("Order event published to Kafka")
	return nil
}
