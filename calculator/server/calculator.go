package main

import (
	"context"
	"log"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
)

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Incoming Add request: %d %d", in.Integer1, in.Integer2)
	total := in.Integer1 + in.Integer2
	return &pb.AddResponse{Value: total}, nil
}

func (s *Server) PrimeFactorization(in *pb.PrimeFactorizationRequest, stream pb.CalculatorService_PrimeFactorizationServer) error {
	log.Printf("PrimeFactorization invoked: %d", in.Integer)

	k := uint32(2)
	n := in.Integer
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeNumber{
				Value: k,
			})
			n = n / k
		} else {
			k = k + 1
		}
	}
	return nil
}
