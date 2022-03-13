package data

import (
	"github.com/Shopify/sarama"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama"

	"dataproducer/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEsClient, NewKafkaProducer, NewDiscovery, NewRegistry, NewOrderRepo)

// Data .
type Data struct {
	// 依赖kafka
	kp sarama.AsyncProducer
	es *elasticsearch.Client
}

// NewData .
func NewData(kp sarama.AsyncProducer, c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	d := &Data{
		kp: kp,
	}
	return d, func() {
		if err := d.kp.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewJobData(es *elasticsearch.Client) (*Data, func(), error) {
	d := &Data{
		es: es,
	}
	return d, func() {
	}, nil
}

func NewAdminData(es *elasticsearch.Client) (*Data, func(), error) {
	d := &Data{
		es: es,
	}
	return d, func() {
	}, nil
}

func NewEsClient(conf *conf.Data) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: conf.Es.Addrs,
		Username:  conf.Es.Username,
		Password:  conf.Es.Password,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return es
}

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	p = otelsarama.WrapAsyncProducer(c, p)
	if err != nil {
		panic(err)
	}
	return p
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistry(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
