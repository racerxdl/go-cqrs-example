package main

import (
	"context"
	"github.com/graphql-go/handler"
	"github.com/racerxdl/go-cqrs-example/gql"
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/protoserver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

var log = logrus.StandardLogger()

func GetReaderClient() protocol.ContactReaderClient {
	log.Infof("Connecting to %s (Reader)", protoserver.LocalConnectReader)
	conn, err := grpc.Dial(protoserver.LocalConnectReader, protoserver.DialOptions)
	if err != nil {
		log.Fatalf("Cannot connect to reader: %s", err)
	}

	return protocol.NewContactReaderClient(conn)
}

func GetWriterClient() protocol.ContactWriterClient {
	log.Infof("Connecting to %s (Writer)", protoserver.LocalConnectWriter)
	conn, err := grpc.Dial(protoserver.LocalConnectWriter, protoserver.DialOptions)
	if err != nil {
		log.Fatalf("Cannot connect to Writer: %s", err)
	}

	return protocol.NewContactWriterClient(conn)
}

func main() {
	cr := GetReaderClient()
	cw := GetWriterClient()

	schema, err := gql.GetSchema()
	if err != nil {
		log.Fatalf("Error reading schema: %s", err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	// Create context and inject Clients
	ctx := context.Background()
	ctx = protocol.InjectContactReaderClientInContext(ctx, cr)
	ctx = protocol.InjectContactWriterClientInContext(ctx, cw)

	// Attach the normal query / mutation handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlHandler.ContextHandler(ctx, w, r)
	})

	log.Println("Listening in :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error listening: %s", err)
	}
}
