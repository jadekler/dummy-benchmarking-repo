package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/jadekler/dummy-benchmarking-repo/.kokoro/grpc"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStorageClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.Read(ctx, &pb.ObjectRead{BucketName: "bucket", ObjectName: "object"})
	if err != nil {
		log.Fatalf("could not read: %v", err)
		fmt.Printf("could not read: %v \n", err)
	}
	fmt.Printf("Read: %v \n", r.Number)
	w, err := c.Write(ctx, &pb.ObjectWrite{BucketName: "bucket", ObjectName: "object", Destination: "destination"})
	if err != nil {
		log.Fatalf("could not write: %v", err)
		fmt.Printf("could not write: %v \n", err)
	}
	fmt.Printf("Write: %v \n", w.Number)
}
