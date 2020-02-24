package main

import (
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/protoserver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var log = logrus.StandardLogger()

func main() {
	lis, err := net.Listen("tcp", protoserver.WriterListen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	server := protoserver.MakeContactWriter()

	protocol.RegisterContactWriterServer(grpcServer, server)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving gRPC: %s", err)
	}
}
