package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	pb "mytest/go-proto-order/ecommerce"
	"net"
)

type Order struct {
	pb.UnimplementedOrderManagementServer
}

func (r *Order) GetOrder(ctx context.Context, str *wrapperspb.StringValue) (*pb.Order, error) {
	return &pb.Order{}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &Order{})
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
