package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello clarissa\n"))
	})

	http.HandleFunc("/bar", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello from bar\n"))
	})

	http.HandleFunc("/some-bucket-name/some-object-name", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("grpc connect to fake server \n"))
		fmt.Print("Here\n")
	})

	fmt.Println("listening on localhost:8080")
	panic(http.ListenAndServe(":8080", nil))
}
