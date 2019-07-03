module github.com/jadekler/dummy-benchmarking-repo

go 1.12

require (
	cloud.google.com/go v0.38.0
	//cloud.google.com/go/storage v0.38.0
	github.com/golang/protobuf v1.3.1
	github.com/googleapis/gax-go v2.0.2+incompatible
	google.golang.org/api v0.7.0
	google.golang.org/grpc v1.21.1
)

replace cloud.google.com/go v0.38.0 => ../google-cloud-go
