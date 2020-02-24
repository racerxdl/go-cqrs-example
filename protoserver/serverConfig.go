package protoserver

import "google.golang.org/grpc"

const (
	// Contact Server Reader DSN
	ReaderListen = ":9998"
	// Contact Server Writer DSN
	WriterListen = ":9999"

	// Local connection DSN for Reader
	LocalConnectReader = "localhost" + ReaderListen

	// Local connection DSN for Writer
	LocalConnectWriter = "localhost" + WriterListen
)

// Dial Options for an client. Defaults to Insecure
var DialOptions = grpc.WithInsecure()
