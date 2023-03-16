package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	pb "mytest/go-proto-order/ecommerce"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)
	order, err := client.GetOrder(context.Background(), &wrapperspb.StringValue{Value: "hello world"})
	if err != nil {
		return
	}
	fmt.Println(order.Description)
}