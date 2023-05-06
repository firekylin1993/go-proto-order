package main

import (
	"context"
	pb "go-proto-order/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net"
)

type Order struct {
	pb.UnimplementedOrderManagementServer
}

func (r *Order) GetOrder(ctx context.Context, str *wrapperspb.StringValue) (*pb.Order, error) {
	return &pb.Order{
		Description: "hello " + str.GetValue(),
	}, nil
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
