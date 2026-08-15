package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("serving on :50051")
	s.Serve(lis)
}
