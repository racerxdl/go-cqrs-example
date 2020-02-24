package main

import (
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/protoserver"
	"github.com/racerxdl/go-cqrs-example/queueManager"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var log = logrus.StandardLogger()

func GetWriterDatabaseClient() protocol.ContactWriterClient {
	log.Infof("Connecting to %s (Database Writer)", protoserver.LocalConnectDatabase)
	conn, err := grpc.Dial(protoserver.LocalConnectDatabase, protoserver.DialOptions)
	if err != nil {
		log.Fatalf("Cannot connect to reader: %s", err)
	}

	return protocol.NewContactWriterClient(conn)
}

func main() {
	q, err := queueManager.MakeMQTTQueueManager("tcp://localhost:1883")
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %s", err)
	}

	qw := protoserver.MakeQueueWriter(q, GetWriterDatabaseClient())

	log.Infof("Queue Writer initialized and waiting")
	qw.Wait()
}
