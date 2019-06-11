package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/jadekler/dummy-benchmarking-repo/.kokoro/grpc"
	"github.com/jadekler/dummy-benchmarking-repo/storage"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Read(ctx context.Context, in *pb.ObjectRead) (*pb.Completed, error) {
	storage.Read("a", in.BucketName, in.ObjectName)
	fmt.Printf("%s\n", "Reading")
	return &pb.Completed{Number: 2, Error: ""}, nil
}

func (s *server) Write(ctx context.Context, in *pb.ObjectWrite) (*pb.Completed, error) {
	storage.Write("a", in.BucketName, in.ObjectName, in.Destination)
	fmt.Printf("%s\n", "Writting")
	return &pb.Completed{Number: 1, Error: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		fmt.Printf("failed to serve: %v", err)
	}
}
