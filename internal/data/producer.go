package data

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"

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

type orderJobRepo struct {
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

// 持久化到 es 中
func (o *orderJobRepo) PersistentSaveOrder(ctx context.Context, order *biz.Order) (*biz.Order, error) {
	// Build the request body.
	var b strings.Builder
	data, _ := json.Marshal(order)
	b.WriteString(string(data))
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      "order",
		DocumentID: fmt.Sprintf("%d", order.OrderId),
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	// Perform the request with the client.
	res, err := req.Do(context.Background(), o.data.es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Infof("[%s] Error indexing document ID=%d", res.Status(), order.OrderId)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Infof("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Infof("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return order, err
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewOrderJobRepo(data *Data, logger log.Logger) biz.OrderJobRepo {
	return &orderJobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
