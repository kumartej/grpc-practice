package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/kgedala/grpc-practice/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true

	opts := []grpc.DialOption{}
	if tls {
		crtFile := "ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(crtFile, "")
		if err != nil {
			log.Fatalf("Loading certs failed: %v\n", creds)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	DoGreetManyTimes(c)
	doLongGreet(c)
	doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 1*time.Second)
}
