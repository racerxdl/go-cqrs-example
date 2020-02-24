package main

import (
	"context"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

type contactManager struct {
	db *Database
}

func MakeContactManager(db *Database) *contactManager {
	return &contactManager{
		db: db,
	}
}

func (cm *contactManager) GetContact(ctx context.Context, contactReference *protocol.ContactReference) (*protocol.ContactRequestResponse, error) {
	c, err := cm.db.GetContact(contactReference)

	response := protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "OK",
	}

	if err != nil {
		response.Status = protocol.RequestResponse_ERROR
		response.Message = err.Error()
	}

	return &protocol.ContactRequestResponse{
		Contact:  c,
		Response: &response,
	}, nil
}

func (cm *contactManager) ListContacts(ctx context.Context, filter *protocol.ListContactsFilter) (*protocol.ContactArrayRequestResponse, error) {
	c, err := cm.db.ListContacts(filter)

	response := protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "OK",
	}

	if err != nil {
		response.Status = protocol.RequestResponse_ERROR
		response.Message = err.Error()
	}

	return &protocol.ContactArrayRequestResponse{
		Contact:  c,
		Response: &response,
	}, nil
}

func (cm *contactManager) AddContact(ctx context.Context, contact *protocol.Contact) (*protocol.RequestResponse, error) {
	err := cm.db.AddContact(contact)

	response := protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "OK",
	}

	if err != nil {
		response.Status = protocol.RequestResponse_ERROR
		response.Message = err.Error()
	}

	return &response, nil
}

func (cm *contactManager) UpdateContact(ctx context.Context, contact *protocol.Contact) (*protocol.RequestResponse, error) {
	err := cm.db.UpdateContact(contact)

	response := protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "OK",
	}

	if err != nil {
		response.Status = protocol.RequestResponse_ERROR
		response.Message = err.Error()
	}

	return &response, nil
}

func (cm *contactManager) DeleteContact(ctx context.Context, contactReference *protocol.ContactReference) (*protocol.RequestResponse, error) {
	err := cm.db.DeleteContact(contactReference)

	response := protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "OK",
	}

	if err != nil {
		response.Status = protocol.RequestResponse_ERROR
		response.Message = err.Error()
	}

	return &response, nil
}
