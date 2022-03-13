package job

import (
	"github.com/Shopify/sarama"
	"github.com/starryrbs/kfan/pkg/job"

	"dataproducer/internal/biz"
	"dataproducer/internal/conf"
)

func NewJobServer(ojr biz.OrderJobRepo, c *conf.Data, kcg sarama.ConsumerGroup) *job.SaramaConsumerJobServer {
	kc := NewConsumeHandler(ojr)
	return job.NewSaramaConsumerJobServer(
		[]string{c.Kafka.Topic},
		kc,
		kcg,
	)
}
