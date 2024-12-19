package producer

import (
	"context"
	"encoding/json"
	"gotinder/util"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaTopic     = "test-topic"
	kafkaBrokerURL = "localhost:9092"
)

func ProducerMessage(message util.RequestMessage) error {
	writer := kafka.Writer{
		Addr:         kafka.TCP(kafkaBrokerURL),
		Topic:        kafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}
	defer writer.Close()
	jsonData, _ := json.Marshal(message)
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte(string(jsonData)),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
