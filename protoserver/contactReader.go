package protoserver

import (
	"context"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

type contactReader struct {
	client protocol.ContactReaderClient
}

func MakeContactReaderProxy(client protocol.ContactReaderClient) protocol.ContactReaderServer {
	return &contactReader{
		client: client,
	}
}

func (cr *contactReader) GetContact(ctx context.Context, contactReference *protocol.ContactReference) (*protocol.ContactRequestResponse, error) {
	return cr.client.GetContact(ctx, contactReference)
}

func (cr *contactReader) ListContacts(ctx context.Context, listContactsFilter *protocol.ListContactsFilter) (*protocol.ContactArrayRequestResponse, error) {
	return cr.client.ListContacts(ctx, listContactsFilter)
}
