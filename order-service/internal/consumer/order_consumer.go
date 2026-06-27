package consumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "order-group",
	})
}

func Consume(reader *kafka.Reader) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("Received: %s\n", string(msg.Value))
	}
}
