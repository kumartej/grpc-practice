package main

import (
	"context"
	"log"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, number int32) {
	log.Printf("doSqrt was invoked with %v\n", number)

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Integer: number,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from Server: %v\n", e.Message())
			log.Printf("Error code from Server: %v\n", e.Code())
		} else {
			log.Fatalf("Uknown gRPC error: %v\n", err)
		}
		return
	}

	log.Printf("SQRT Response from Server: %v", res.Value)
}
