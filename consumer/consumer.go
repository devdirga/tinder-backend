package consumer

import (
	"context"
	"encoding/json"
	"gotinder/util"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaTopic     = "test-topic"
	kafkaBrokerURL = "localhost:9092"
)

func ConsumeMessages() error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBrokerURL},
		Topic:   kafkaTopic,
		GroupID: "example-group",
	})
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var mail util.RequestMessage
		err = json.Unmarshal([]byte(msg.Value), &mail)
		if err != nil {
			return err
		}

		util.SendMail(map[string]interface{}{
			"to":      mail.To,
			"subject": mail.Subject,
			"message": mail.Message,
		})
	}
}
