package storage

import (
	"time"
)

func Read(foo, bucketName string, objectName string) {
	time.Sleep(2000 * time.Millisecond)
}

func Write(foo, bucketName string, objectName string, destination string){
	time.Sleep(5000 * time.Millisecond)
}