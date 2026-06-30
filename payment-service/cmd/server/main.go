package main

import (
	"github.com/madhurima877/payment-service/internal/consumer"
	"github.com/madhurima877/payment-service/internal/producer"
)

func main() {
	writer := producer.NewWriter()
	defer writer.Close()
	reader := consumer.NewPaymentReader()
	defer reader.Close()
	consumer.ConsumePayment(reader, writer)
}
