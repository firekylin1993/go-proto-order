package main

import (
	"context"
	"fmt"
	pb "go-proto-order/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		fmt.Println(err.Error())
		return
	}
	fmt.Println(order.Description)
}
