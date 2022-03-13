package data

import (
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"

	"dataproducer/internal/conf"
)

var JobProviderSet = wire.NewSet(NewJobData, NewKafkaConsumeGroup, NewOrderJobRepo, NewEsClient)

func NewKafkaConsumeGroup(c *conf.Data) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroup(c.Kafka.Addrs, c.Kafka.GroupId, sarama.NewConfig())
}
