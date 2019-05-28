package main

import (
	"context"
	"fmt"
	"io"
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
	// doSum(cc)
	c := calculatorpb.NewCalculatorServiceClient(cc)
	doServerStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecompoisition Server Streaming RPC ...!")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 123634634,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happend: %v", err)
		}
		fmt.Printf("Prime factor is: %v\n", res.GetPrimeFactor())
	}
}
