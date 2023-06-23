package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kgedala/grpc-practice/greet/proto"
)

func DoGreetManyTimes(c pb.GreetServiceClient) {
	log.Printf("DoGreetManyTimes was invoked")

	client, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		FirstName: "Kumar",
	})
	if err != nil {
		log.Fatalf("Error Querying GreetManyTimes: %v", err)
	}

	for {
		res, err := client.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading from stream: %v\n", err)
		}
		log.Printf("Received: %s", res)
	}
}
