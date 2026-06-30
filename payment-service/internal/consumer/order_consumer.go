package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/madhurima877/payment-service/internal/model"
	"github.com/madhurima877/payment-service/internal/producer"
	"github.com/segmentio/kafka-go"
)

func NewPaymentReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "payment-group",
	})
}

func ConsumePayment(reader *kafka.Reader, writer *kafka.Writer) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		var order model.Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Println(err)
			continue
		}

		log.Printf(
			"Processing payment for Order %d\n",
			order.ID,
		)

		payment := model.Payment{
			OrderID: order.ID,
			Status:  "COMPLETED",
		}
		paymentEvent, err := json.Marshal(payment)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := producer.Publish(writer, paymentEvent); err != nil {
			log.Println(err)
			continue
		}
		log.Println("PaymentCompleted event published")
	}
}
