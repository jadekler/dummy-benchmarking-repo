package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/storage"
	pb "github.com/jadekler/dummy-benchmarking-repo/storage-go/proto"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	c *storage.Client
}

func (s *server) Read(ctx context.Context, in *pb.ObjectRead) (*pb.EmptyResponse, error) {
	b := s.c.Bucket(in.GetBucketName())
	o := b.Object(in.GetObjectName())
	r, err := o.NewReader(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	var ba []byte
	fmt.Print("Is reading\n")
	if _, err = r.Read(ba); err != nil {
		log.Fatal(err)
	}
	/*for {
		var b []byte
		fmt.Print("Is reading")
		_, err := r.Read(b)
		fmt.Print("Finish")
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Do nothing with b.
	}*/
	return &pb.EmptyResponse{}, nil
}

func (s *server) Write(ctx context.Context, in *pb.ObjectWrite) (*pb.EmptyResponse, error) {
	// (do something similar as Read)
	return &pb.EmptyResponse{}, nil
}

func main() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx, option.WithEndpoint("http://localhost:8080/"))
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterStorageServer(s, &server{
		c: c,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
