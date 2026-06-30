package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/madhurima877/analytics-service/internal/model"
	"github.com/segmentio/kafka-go"
)

func NewAnalyticsReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "payments",
		GroupID: "analytics-group",
	})

}

func ConsumeAnalytics(reader *kafka.Reader) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}
		var payment model.Payment
		if err := json.Unmarshal(msg.Value, &payment); err != nil {
			log.Println(err)
			continue
		}
		log.Printf(
			"Analytics: Payment completed for Order %d\n",
			payment.OrderID,
		)
	}
}
