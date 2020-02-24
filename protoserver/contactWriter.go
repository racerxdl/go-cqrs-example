package protoserver

import (
	"context"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

type contactWriter struct {
}

func MakeContactWriter() protocol.ContactWriterServer {
	return nil
}

func (cw *contactWriter) AddContact(context.Context, *protocol.Contact) (*protocol.RequestResponse, error) {
	return nil, nil
}

func (cw *contactWriter) UpdateContact(context.Context, *protocol.Contact) (*protocol.RequestResponse, error) {
	return nil, nil

}

func (cw *contactWriter) DeleteContact(context.Context, *protocol.ContactReference) (*protocol.RequestResponse, error) {
	return nil, nil

}
