package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gunplan.top/SCI/protocol"
	"gunplan.top/SCI/server/service"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8028")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	s := grpc.NewServer()
	protocol.RegisterVerifyServer(s, &service.VerifyImpl{})
	protocol.RegisterPageServer(s, &service.PageImpl{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
