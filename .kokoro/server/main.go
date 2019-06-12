package main

import (
	"context"
	"log"
	"net"

	pb "github.com/jadekler/dummy-benchmarking-repo/.kokoro/grpc"
	"github.com/jadekler/dummy-benchmarking-repo/storage"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (s *server) Read(ctx context.Context, in *pb.ObjectRead) (*pb.EmptyResponse, error) {
	storage.Read(in.BucketName, in.ObjectName)
	return &pb.EmptyResponse{}, nil
}

func (s *server) Write(ctx context.Context, in *pb.ObjectWrite) (*pb.EmptyResponse, error) {
	storage.Write(in.BucketName, in.ObjectName, in.Destination)
	return &pb.EmptyResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
