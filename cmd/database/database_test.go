package main

import (
	"github.com/racerxdl/go-cqrs-example/protocol"
	"testing"
	"time"
)

func TestMakeMemoryDatabase(t *testing.T) {
	_, err := MakeMemoryDatabase()

	if err != nil {
		t.Fatalf("MakeMemoryDatabase should not throw errors: %s", err)
	}
}

func TestDatabase_AddContact(t *testing.T) {
	db, _ := MakeMemoryDatabase()
	err := db.AddContact(&protocol.Contact{
		Name: "Test User",
	})

	if err != nil {
		t.Fatalf("error adding user: %s", err)
	}

	contacts, _ := db.ListContacts(&protocol.ListContactsFilter{Count: 1})

	if len(contacts) == 0 {
		t.Fatalf("user not added")
	}

	c := contacts[0]
	if c.Name != "Test User" {
		t.Fatalf("Expected user name to be \"Test User\" but got \"%s\"", c.Name)
	}
}

func TestDatabase_UpdateContact(t *testing.T) {
	db, _ := MakeMemoryDatabase()
	_ = db.AddContact(&protocol.Contact{
		Name: "Test User",
	})

	time.Sleep(time.Second) // Increment Timestamp

	contacts, _ := db.ListContacts(&protocol.ListContactsFilter{Count: 1})

	c := contacts[0]

	c.Name = "Updated Test User"

	err := db.UpdateContact(c)

	if err != nil {
		t.Fatalf("Failed to update user: %s", err)
	}

	updatedContact, err := db.GetContact(&protocol.ContactReference{Id: c.Id})

	if updatedContact.Name != c.Name {
		t.Fatalf("Expected user name change to %s got %s", c.Name, updatedContact.Name)
	}

	if updatedContact.LastUpdated.Seconds == c.LastUpdated.Seconds {
		t.Fatalf("Expected timestamp to be updated")
	}
}

func TestDatabase_DeleteContact(t *testing.T) {
	db, _ := MakeMemoryDatabase()
	_ = db.AddContact(&protocol.Contact{
		Name: "Test User",
	})

	contacts, _ := db.ListContacts(&protocol.ListContactsFilter{Count: 1})

	c := contacts[0]

	err := db.DeleteContact(&protocol.ContactReference{Id: c.Id})

	if err != nil {
		t.Fatalf("Failed to update user: %s", err)
	}

	_, err = db.GetContact(&protocol.ContactReference{Id: c.Id})

	if err == nil {
		t.Fatalf("Expected error to be not found")
	}
}
