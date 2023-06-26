package main

import (
	"log"
	"net"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	tls := true

	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to load certs: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	conn := grpc.NewServer(opts...)

	pb.RegisterCalculatorServiceServer(conn, &Server{})

	err = conn.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
