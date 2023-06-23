package main

import (
	"fmt"
	"log"

	pb "github.com/kgedala/grpc-practice/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Invoked GreetmanyTimes: %v", in)

	for i := 1; i < 10; i++ {
		res := fmt.Sprintf("Hello %s , number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
