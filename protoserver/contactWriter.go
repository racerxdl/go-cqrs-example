package protoserver

import (
	"context"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

type contactWriter struct {
	client protocol.ContactWriterClient
}

func MakeContactWriterProxy(client protocol.ContactWriterClient) protocol.ContactWriterServer {
	return &contactWriter{
		client: client,
	}
}

func (cw *contactWriter) AddContact(ctx context.Context, contact *protocol.Contact) (*protocol.RequestResponse, error) {
	return cw.client.AddContact(ctx, contact)
}

func (cw *contactWriter) UpdateContact(ctx context.Context, contact *protocol.Contact) (*protocol.RequestResponse, error) {
	return cw.client.UpdateContact(ctx, contact)
}

func (cw *contactWriter) DeleteContact(ctx context.Context, contactReference *protocol.ContactReference) (*protocol.RequestResponse, error) {
	return cw.client.DeleteContact(ctx, contactReference)
}
