package main

import (
	"google.golang.org/grpc"
	"gunplan.top/SCI/protocol"
	"gunplan.top/SCI/server/service"
)

func regisite(s *grpc.Server) {
	protocol.RegisterVerifyServer(s, &service.VerifyImpl{})
	protocol.RegisterPageServer(s, &service.PageImpl{})
}

