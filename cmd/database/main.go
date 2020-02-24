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
	db, err := MakeMemoryDatabase()

	if err != nil {
		log.Fatalf("Error initializing database: %s", err)
	}

	contactManager := MakeContactManager(db)

	lis, err := net.Listen("tcp", protoserver.DatabaseListen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	protocol.RegisterContactWriterServer(grpcServer, contactManager)
	protocol.RegisterContactReaderServer(grpcServer, contactManager)

	log.Infof("Listening in %s", protoserver.DatabaseListen)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error serving gRPC: %s", err)
	}
}
