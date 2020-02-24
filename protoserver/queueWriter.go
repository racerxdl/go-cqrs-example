package protoserver

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/racerxdl/go-cqrs-example/queueManager"
	"github.com/sirupsen/logrus"
	"reflect"
	"time"
)

var log = logrus.StandardLogger()

const (
	writeDelay = time.Second * 5
)

type queueWriter struct {
	q queueManager.QueueManager
	w protocol.ContactWriterClient
}

func MakeQueueWriter(queue queueManager.QueueManager, databaseWriter protocol.ContactWriterClient) *queueWriter {
	qw := &queueWriter{
		q: queue,
		w: databaseWriter,
	}

	qw.init()

	return qw
}

func (qw *queueWriter) init() {
	log.Infof("Subscribing to topic %s", addContactTopic)
	qw.q.Subscribe(addContactTopic, qw.addContact)
	log.Infof("Subscribing to topic %s", deleteContactTopic)
	qw.q.Subscribe(deleteContactTopic, qw.deleteContact)
	log.Infof("Subscribing to topic %s", updateContactTopic)
	qw.q.Subscribe(updateContactTopic, qw.updateContact)
}

func (qw *queueWriter) addContact(data interface{}) {
	if dataBytes, ok := data.([]byte); ok {
		contact := protocol.Contact{}
		err := proto.Unmarshal(dataBytes, &contact)
		if err != nil {
			log.Errorf("Received invalid addContact data: %s", err)
			return
		}
		log.Infof("addContact(%s): received", contact.Name)
		time.Sleep(writeDelay)
		resp, err := qw.w.AddContact(context.Background(), &contact)
		if err != nil {
			log.Errorf("Error writing to database: %s", err)
			return
		}
		if resp.Status != protocol.RequestResponse_OK {
			log.Errorf("Error writing to database: %s", resp.Message)
			return
		}
		log.Infof("addContact(%s): done", contact.Name)
	} else {
		log.Errorf("Invalid type received. Expected bytes got %s", reflect.TypeOf(data).Name())
	}
}

func (qw *queueWriter) deleteContact(data interface{}) {
	if dataBytes, ok := data.([]byte); ok {
		contact := protocol.ContactReference{}
		err := proto.Unmarshal(dataBytes, &contact)
		if err != nil {
			log.Errorf("Received invalid addContact data: %s", err)
			return
		}
		log.Infof("deleteContact(%s): received", contact.Id)
		time.Sleep(writeDelay)
		resp, err := qw.w.DeleteContact(context.Background(), &contact)
		if err != nil {
			log.Errorf("Error writing to database: %s", err)
			return
		}
		if resp.Status != protocol.RequestResponse_OK {
			log.Errorf("Error writing to database: %s", resp.Message)
			return
		}
		log.Infof("deleteContact(%s): done", contact.Id)
	} else {
		log.Errorf("Invalid type received. Expected bytes got %s", reflect.TypeOf(data).Name())
	}
}

func (qw *queueWriter) updateContact(data interface{}) {
	if dataBytes, ok := data.([]byte); ok {
		contact := protocol.Contact{}
		err := proto.Unmarshal(dataBytes, &contact)
		if err != nil {
			log.Errorf("Received invalid addContact data: %s", err)
			return
		}
		log.Infof("updateContact(%s, %s): received", contact.Id, contact.Name)
		time.Sleep(writeDelay)
		resp, err := qw.w.UpdateContact(context.Background(), &contact)
		if err != nil {
			log.Errorf("Error writing to database: %s", err)
			return
		}
		if resp.Status != protocol.RequestResponse_OK {
			log.Errorf("Error writing to database: %s", resp.Message)
			return
		}
		log.Infof("updateContact(%s, %s): done", contact.Id, contact.Name)
	} else {
		log.Errorf("Invalid type received. Expected bytes got %s", reflect.TypeOf(data).Name())
	}
}

func (qw *queueWriter) Wait() {
	select {
	// Do nothing for now
	}
}
