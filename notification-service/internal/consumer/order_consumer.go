package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/madhurima877/notification-service/internal/model"
	"github.com/segmentio/kafka-go"
)

func NewNotificationReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "notification-group",
	})
}

func ConsumeNotification(reader *kafka.Reader) {
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
			"Sending notification: Order %d for %s created\n",
			order.ID,
			order.Product,
		)
	}
}
