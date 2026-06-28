package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/madhurima877/order-service/internal/config"
	"github.com/madhurima877/order-service/internal/handler"
	"github.com/madhurima877/order-service/internal/kafka"
	"github.com/madhurima877/order-service/internal/repository"
	"github.com/madhurima877/order-service/internal/service"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connected")
	writer := kafka.NewWriter()
	defer writer.Close()

	repo := repository.NewOrderRepository(db)
	svc := service.NewOrderService(repo, writer)
	orderHandler := handler.NewOrderHandler(svc)

	fmt.Println("Order Service starting")
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/orders", orderHandler.Create)
	http.ListenAndServe(":8080", nil)
}
