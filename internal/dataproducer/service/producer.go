package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "dataproducer/api/producer"
	"dataproducer/internal/biz"
)

type OrderService struct {
	pb.UnimplementedProducerServer
	oc  *biz.OrderCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc:  oc,
		log: log.NewHelper(log.With(logger, "service/interface"))}
}

func (os *OrderService) SaveOrder(ctx context.Context, req *pb.SaveOrderRequest) (*pb.SaveOrderReply, error) {
	order := new(biz.Order)
	order.OrderId = req.Order.OrderId
	order.PayId = req.Order.PayId
	order.Uid = req.Order.Uid
	order.Price = req.Order.Price
	order.CreateTime = req.Order.CreateTime
	order.UpdateTime = req.Order.UpdateTime
	err := os.oc.Repo.SaveOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &pb.SaveOrderReply{}, nil

}
