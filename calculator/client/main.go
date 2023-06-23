package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)
	res, err := client.Add(context.Background(), &pb.AddRequest{Integer1: 12, Integer2: 35})
	if err != nil {
		log.Fatalf("Error querying add: %v", err)
	}
	log.Printf("Result from add: %d", res.Value)

	cln, err := client.PrimeFactorization(context.Background(), &pb.PrimeFactorizationRequest{
		Integer: 392,
	})
	if err != nil {
		log.Fatalf("Error querying :%v", nil)
	}

	for {
		res, err := cln.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error receiving from server: %v", err)
		}

		log.Printf("%d ", res.Value)
	}
	log.Printf("\n")

}
