package protoserver

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/queueManager"
	"google.golang.org/grpc"
)

type mqttWriterClient struct {
	q queueManager.QueueManager
}

func MakeMQTTWriterClient(queue queueManager.QueueManager) protocol.ContactWriterClient {
	return &mqttWriterClient{
		q: queue,
	}
}

func (mwc *mqttWriterClient) AddContact(ctx context.Context, in *protocol.Contact, opts ...grpc.CallOption) (*protocol.RequestResponse, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	err = mwc.q.Publish(addContactTopic, data)

	if err != nil {
		return nil, err
	}

	return &protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "Queued to add",
	}, nil
}

func (mwc *mqttWriterClient) UpdateContact(ctx context.Context, in *protocol.Contact, opts ...grpc.CallOption) (*protocol.RequestResponse, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	err = mwc.q.Publish(updateContactTopic, data)

	if err != nil {
		return nil, err
	}

	return &protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "Queued to update",
	}, nil
}

func (mwc *mqttWriterClient) DeleteContact(ctx context.Context, in *protocol.ContactReference, opts ...grpc.CallOption) (*protocol.RequestResponse, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}

	err = mwc.q.Publish(deleteContactTopic, data)

	if err != nil {
		return nil, err
	}

	return &protocol.RequestResponse{
		Status:  protocol.RequestResponse_OK,
		Message: "Queued to delete",
	}, nil
}
