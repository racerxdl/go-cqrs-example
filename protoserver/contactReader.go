package protoserver

import (
	"context"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

type contactReader struct {
}

func MakeContactReader() protocol.ContactReaderServer {
	return nil
}

func (cr *contactReader) GetContact(context.Context, *protocol.ContactReference) (*protocol.ContactRequestResponse, error) {
	return nil, nil

}

func (cr *contactReader) ListContacts(context.Context, *protocol.ListContactsFilter) (*protocol.ContactArrayRequestResponse, error) {
	return nil, nil

}
