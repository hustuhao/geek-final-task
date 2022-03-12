package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Order struct {
	OrderId    int64
	Uid        string
	PayId      int64
	Price      int64
	CreateTime int64
	UpdateTime int64
}

type OrderRepo interface {
	SaveOrder(ctx context.Context, order *Order) error
}

type OrderCase struct {
	Repo OrderRepo
	log  *log.Helper
}

func NewOrderCase(repo OrderRepo, logger log.Logger) *OrderCase {
	return &OrderCase{Repo: repo, log: log.NewHelper(log.With(logger, "module", "order"))}
}
