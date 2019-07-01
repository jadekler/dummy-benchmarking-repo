package storage

import (
	"context"
	"log"
	"time"
	"cloud.google.com/go/storage"
)

func Read(ctx context.Context, bucketName string, objectName string) {
	var out []byte
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		log.Fatalf("failed to create reader: %v", err)
	}
	/*data, err := rc.Read(out)
	if err != nil {
		log.Fatalf("failed to read: %v", err)
	}*/
	if _, err := rc.Read(out); err != nil {
		log.Fatalf("failed to read: %v", err)
	}
	//time.Sleep(2000 * time.Millisecond)
}

func Write(ctx context.Context, bucketName string, objectName string, destination string) {
	time.Sleep(5000 * time.Millisecond)
}
