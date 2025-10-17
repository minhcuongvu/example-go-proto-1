package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "example.com/calculator/pb_gen"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received Add request: a=%d, b=%d", req.A, req.B)
	return &pb.AddResponse{Result: req.A + req.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &server{})

	log.Println("Calculator server listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
