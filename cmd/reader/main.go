package main

import (
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/protoserver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var log = logrus.StandardLogger()

func GetReaderDatabaseClient() protocol.ContactReaderClient {
	log.Infof("Connecting to %s (Database Reader)", protoserver.LocalConnectDatabase)
	conn, err := grpc.Dial(protoserver.LocalConnectDatabase, protoserver.DialOptions)
	if err != nil {
		log.Fatalf("Cannot connect to reader: %s", err)
	}

	return protocol.NewContactReaderClient(conn)
}

func main() {
	lis, err := net.Listen("tcp", protoserver.ReaderListen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	server := protoserver.MakeContactReaderProxy(GetReaderDatabaseClient())

	protocol.RegisterContactReaderServer(grpcServer, server)

	log.Infof("Listening in %s", protoserver.ReaderListen)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving gRPC: %s", err)
	}
}
