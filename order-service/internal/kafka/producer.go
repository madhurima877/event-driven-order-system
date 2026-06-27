package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "orders",
	}
}

func Publish(writer *kafka.Writer, message []byte) error {
	return writer.WriteMessages(context.Background(), kafka.Message{Value: message})
}
