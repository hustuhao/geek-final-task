package service

import (
	"context"

	pb "dataproducer/api/admin"
)

type AdminService struct {
	pb.UnimplementedAdminServer
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) SearchOrder(ctx context.Context, req *pb.SearchOrderRequest) (*pb.SearchOrderReply, error) {
	// TODO 查找满足要求的订单

	return &pb.SearchOrderReply{}, nil
}
