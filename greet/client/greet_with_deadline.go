package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, t time.Duration) {
	log.Printf("doGreetWithDeadline was invoked with timeout: %v\n", t)
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Kumar",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.Canceled {
				log.Printf("Timeout Exceeded, Request cancelled from Server")
				return
			} else {
				log.Fatalf("Error message from Server: %v\n", e.Message())
				log.Fatalf("Error code from Server: %v\n", e.Code())
			}
		} else {
			log.Fatalf("Non gRPC Error from Server: %v\n", err)
		}
	}

	log.Printf("Response from server: %v", res.Result)
}
