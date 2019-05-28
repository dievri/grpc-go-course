package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dievri/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello from calculator client!")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer cc.Close()
	doSum(cc)

}

func doSum(cc *grpc.ClientConn) {
	c := calculatorpb.NewCalculatorServiceClient(cc)
	req := &calculatorpb.SumRequest{
		FirstNumber:  1000,
		SecondNumber: 50,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error executing SUM RPC: %v", err)
	}

	fmt.Println(res.GetSumResult())
}
