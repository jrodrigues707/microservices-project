package main

import (
	"fmt"
	"encoding/json"
	"time"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message struct {
    TransactionID int     `json:"transaction_id"`
    UserID        int     `json:"user_id"`
    ProductID     int     `json:"product_id"`
    Quantity      int     `json:"quantity"`
    Price         float64 `json:"price"`
    Timestamp     string  `json:"timestamp"`
}

func main() {
	time.Sleep(20 * time.Second)
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "my-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		return
	}

	c.SubscribeTopics([]string{"my-topic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			var message Message
			json.Unmarshal(msg.Value, &message)
			fmt.Printf("Consumed transaction: %+v (on %s)\n", message, msg.TopicPartition)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}