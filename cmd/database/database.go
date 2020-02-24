package main

import (
	"context"
	"github.com/256dpi/lungo"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/racerxdl/go-cqrs-example/protocol"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Database struct {
	client lungo.IClient
	engine *lungo.Engine
}

func MakeMemoryDatabase() (*Database, error) {
	// prepare options
	opts := lungo.Options{
		Store: lungo.NewMemoryStore(),
	}

	// open database
	client, engine, err := lungo.Open(nil, opts)
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
		engine: engine,
	}, nil
}

func (d *Database) Close() {
	d.engine.Close()
}

func GetNowTimestamp() *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: time.Now().Unix(),
	}
}

func (d *Database) AddContact(contact *protocol.Contact) error {
	c := d.client.Database("default").Collection("contacts")
	contact.Id = uuid.NewV4().String()
	contact.LastUpdated = GetNowTimestamp()
	_, err := c.InsertOne(context.Background(), contact)
	return err
}

func (d *Database) UpdateContact(contact *protocol.Contact) error {
	c := d.client.Database("default").Collection("contacts")

	_, err := c.UpdateOne(context.Background(), bson.D{
		{"id", contact.Id},
	}, bson.D{
		{"$set", bson.D{
			{"name", contact.Name},
			{"lastUpdated", GetNowTimestamp()},
		}},
	})

	return err
}

func (d *Database) DeleteContact(contact *protocol.ContactReference) error {
	c := d.client.Database("default").Collection("contacts")

	_, err := c.DeleteOne(context.Background(), bson.D{
		{"id", contact.Id},
	})

	return err
}

func (d *Database) GetContact(contactReference *protocol.ContactReference) (*protocol.Contact, error) {
	c := d.client.Database("default").Collection("contacts")

	contact := protocol.Contact{}

	err := c.FindOne(context.Background(), bson.D{
		{"id", contactReference.Id},
	}).Decode(&contact)

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (d *Database) ListContacts(filter *protocol.ListContactsFilter) ([]*protocol.Contact, error) {
	c := d.client.Database("default").Collection("contacts")
	cursor, err := c.Find(context.Background(), bson.D{}, &options.FindOptions{
		// Max: int(contact.Count), // Lungo does not support max
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	contacts := make([]*protocol.Contact, 0)

	for cursor.Next(context.Background()) {
		contact := protocol.Contact{}
		err = cursor.Decode(&contact)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, &contact)

		if len(contacts) == int(filter.Count) {
			break
		}
	}

	return contacts, nil
}
