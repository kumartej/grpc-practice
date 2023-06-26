package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/kgedala/grpc-practice/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function invoked")

	var result float32 = 0
	var count int = 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error during the request: %v", err)
			return err
		}

		result = result + float32(res.Integer)
		count = count + 1

	}

	err := stream.SendAndClose(&pb.DoubleResponse{
		Value: result / float32(count),
	})
	if err != nil {
		log.Fatalf("Error while sending response: %v\n", err)
		return err
	}

	return nil
}

func (s *Server) CurrentMax(stream pb.CalculatorService_CurrentMaxServer) error {
	log.Println("CurrentMax was invoked")

	var max uint32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading from client: %v\n", err)
		}

		if req.Integer > max {
			max = req.Integer
			err = stream.Send(&pb.IntegerResponse{
				Value: req.Integer,
			})

			if err != nil {
				log.Fatalf("Error while sending to client: %v", err)
			}
		}

	}
}

func (s *Server) Sqrt(c context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt invoked with %v\n", in)

	num := in.Integer
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", num),
		)
	}

	return &pb.SqrtResponse{
		Value: math.Sqrt(float64(num)),
	}, nil
}
