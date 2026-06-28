package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/madhurima877/order-service/internal/model"
	"github.com/segmentio/kafka-go"
)

func NewInventoryReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "inventory-group",
	})
}
func ConsumerInventory(reader *kafka.Reader) {
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
		log.Printf("Reducing inventory for %s by %d\n",
			order.Product,
			order.Quantity)
	}
}
