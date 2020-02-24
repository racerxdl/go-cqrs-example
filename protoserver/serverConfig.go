package protoserver

import "google.golang.org/grpc"

const (
	// Contact Server Reader DSN
	ReaderListen = ":9998"
	// Contact Server Writer DSN
	WriterListen = ":9999"

	// Contact Server Reader DSN
	DatabaseListen = ":8888"

	// Local connection DSN for Reader
	LocalConnectReader = "localhost" + ReaderListen

	// Local connection DSN for Writer
	LocalConnectWriter = "localhost" + WriterListen

	// Local connection DSN for Writer
	LocalConnectDatabase = "localhost" + DatabaseListen
)

// Dial Options for an client. Defaults to Insecure
var DialOptions = grpc.WithInsecure()
