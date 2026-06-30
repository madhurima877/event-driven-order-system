package main

import "github.com/madhurima877/analytics-service/internal/consumer"

func main() {
	reader := consumer.NewAnalyticsReader()
	defer reader.Close()
	consumer.ConsumeAnalytics(reader)
}

// A client sends a POST /orders request to the Order Service. The Order Service validates the request, saves the order in PostgreSQL, and publishes an OrderCreated event to the orders Kafka topic. Inventory Service, Notification Service, and Payment Service each belong to different consumer groups, so they all receive the same event independently. Inventory updates stock, Notification sends a notification, and Payment processes the payment. After processing, the Payment Service publishes a new PaymentCompleted event to the payments topic. The Analytics Service subscribes to the payments topic and receives that event, where it logs or processes analytics for the completed payment.
