package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kgedala/grpc-practice/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Kumar"},
		{FirstName: "Abhijeet"},
		{FirstName: "Vaibhav"},
	}

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
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
				log.Printf("Error while receiving: %v\b", err)
				break
			}

			log.Printf("GreetEveryone Received %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
