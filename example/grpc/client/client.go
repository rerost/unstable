package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	api_pb "github.com/rerost/unstable/example/grpc/server/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := api_pb.NewSampleServiceClient(conn)

	ctx := context.Background()
	r, err := c.GetSample(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	log.Fatalf("Response: %v", r.Message)
}
