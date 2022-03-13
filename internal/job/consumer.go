package job

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"

	"dataproducer/internal/biz"
)

type consumeHandler struct {
	ready chan bool
	ojr   biz.OrderJobRepo
}

func NewConsumeHandler(hc biz.OrderJobRepo) sarama.ConsumerGroupHandler {
	return &consumeHandler{ready: make(chan bool), ojr: hc}
}

func (c *consumeHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *consumeHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		ctx := context.Background()
		var order biz.Order

		err := json.Unmarshal(message.Value, &order)
		if err != nil {
			return err
		}
		_, err = c.ojr.PersistentSaveOrder(ctx, &order)
		if err != nil {
			return err
		}
		session.MarkMessage(message, "")
	}
	return nil
}
