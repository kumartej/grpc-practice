package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kumar"},
		{FirstName: "Abhijeet"},
		{FirstName: "Vaibhav"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending res: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Long Greet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
