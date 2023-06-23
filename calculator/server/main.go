package main

import (
	"log"
	"net"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Server Listen failed: %v\n", err)
	}
	log.Println("Listening on 50052")
	conn := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(conn, &Server{})

	err = conn.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
