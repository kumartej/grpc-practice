package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
)

func doCurrentMax(c pb.CalculatorServiceClient) {
	log.Println("doCurrentMax was invoked")

	reqs := []*pb.IntegerRequest{
		{Integer: 1},
		{Integer: 5},
		{Integer: 3},
		{Integer: 6},
		{Integer: 2},
		{Integer: 20},
	}

	stream, err := c.CurrentMax(context.Background())
	if err != nil {
		log.Fatalf("Error sending to client")
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error reading from server: %v\n", err)
				break
			}

			log.Printf("Received Max %v\n", res.Value)
		}
		close(waitc)
	}()

	<-waitc
}
