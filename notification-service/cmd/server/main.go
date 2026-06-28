package main

import "github.com/madhurima877/notification-service/internal/consumer"

func main() {
	reader := consumer.NewNotificationReader()
	defer reader.Close()
	consumer.ConsumeNotification(reader)
}
