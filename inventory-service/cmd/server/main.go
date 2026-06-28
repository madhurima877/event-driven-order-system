package main

import "github.com/madhurima877/inventory-service/internal/consumer"

func main() {
	reader := consumer.NewInventoryReader()
	defer reader.Close()
	consumer.ConsumeInventory(reader)
}
