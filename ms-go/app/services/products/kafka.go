package products

import (
	"context"
	"encoding/json"
	"fmt"
	"ms-go/app/models"

	"github.com/segmentio/kafka-go"
)

func sendMessageToKafka(product models.Product) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP("kafka:29092"),
		Topic:    "go-to-rails",
		Balancer: &kafka.LeastBytes{},
	}

	message, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to serialize product: %v", err)
	}

	err = writer.WriteMessages(context.TODO(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %v", err)
	}

	return nil
}