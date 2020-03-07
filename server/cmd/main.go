package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8028")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	s := grpc.NewServer()
	regisite(s)
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

