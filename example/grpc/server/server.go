package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	api_pb "github.com/rerost/unstable/example/grpc/server/api"
	"github.com/rerost/unstable/ugrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetSample(ctx context.Context, req *empty.Empty) (*api_pb.GetSampleResponse, error) {
	return &api_pb.GetSampleResponse{Message: "Sample"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(ugrpc.UnstableUnaryServerInterceptor()),
	)
	api_pb.RegisterSampleServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
