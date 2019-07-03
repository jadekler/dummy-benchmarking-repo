package main

import (
	"context"
	"fmt"

	pb "github.com/jadekler/dummy-benchmarking-repo/storage-go/proto"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	cl := pb.NewStorageClient(conn)
	c := context.Background()
	fmt.Print("Start to read\n")
	r, err := cl.Read(c, &pb.ObjectRead{BucketName: "some-bucket-name", ObjectName: "some-object-name"})
	fmt.Print("Finish\n")
	if err != nil {
		fmt.Printf("could not read: %v", err)
	}
	fmt.Printf("Received: %s\n", r)
}
