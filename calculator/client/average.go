package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	var numbers = [5]uint32{2, 4, 6, 8, 10}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while send integers: %v", err)
	}

	for _, num := range numbers {
		log.Printf("Sending Integer: %v\n", num)
		stream.Send(&pb.IntegerRequest{
			Integer: num,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while closing: %v", err)
	}

	log.Printf("Received Average: %v", res.Value)
}
