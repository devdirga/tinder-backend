package producer

import (
	"context"
	"encoding/json"
	"gotinder/config"
	"gotinder/util"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProducerMessage(message util.RequestMessage) error {
	writer := kafka.Writer{
		Addr:         kafka.TCP(config.GetConf().KafkaUrl),
		Topic:        config.GetConf().KafkaTopic,
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
