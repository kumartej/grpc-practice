package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(c context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked: %v\n", in)

	for i := 1; i < 3; i++ {
		if c.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request")
			return nil, status.Errorf(
				codes.Canceled,
				fmt.Sprintf("Request Cancelled"),
			)
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
