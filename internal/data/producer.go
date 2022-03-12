package data

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"
	"go.opentelemetry.io/otel"

	"dataproducer/internal/biz"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

// SaveOrder 写入订单数据到 kafka 中
func (a *orderRepo) SaveOrder(ctx context.Context, order *biz.Order) error {
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// Create root span
	tr := otel.Tracer("producer")
	ctx, span := tr.Start(ctx, "produce message")
	defer span.End()
	msg := sarama.ProducerMessage{
		Topic: "order",
		Value: sarama.ByteEncoder(b),
	}

	otel.GetTextMapPropagator().Inject(ctx, otelsarama.NewProducerMessageCarrier(&msg))
	// 向kafka写入订单数据
	a.data.kp.Input() <- &msg
	return nil
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
