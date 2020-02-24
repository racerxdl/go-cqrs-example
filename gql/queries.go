package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

var rootQuery = graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"contact": &graphql.Field{
			Type: contact,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "ID of the Contact you want to fetch",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cr := protocol.ContactReaderClientFromContext(p.Context)
				if cr == nil {
					return nil, fmt.Errorf("no contact reader client")
				}

				resp, err := cr.GetContact(p.Context, &protocol.ContactReference{
					Id: p.Args["id"].(string),
				})

				if err != nil {
					return nil, err
				}

				if resp.Response.Status != protocol.RequestResponse_OK {
					return nil, fmt.Errorf(resp.Response.Message)
				}

				return resp.Contact, nil
			},
		},
		"listContacts": &graphql.Field{
			Type: graphql.NewList(contact),
			Args: graphql.FieldConfigArgument{
				"count": &graphql.ArgumentConfig{
					Description: "First N items of the list",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cr := protocol.ContactReaderClientFromContext(p.Context)
				if cr == nil {
					return nil, fmt.Errorf("no contact reader client")
				}

				resp, err := cr.ListContacts(p.Context, &protocol.ListContactsFilter{
					Count: int32(p.Args["count"].(int)),
				})

				if err != nil {
					return nil, err
				}

				if resp.Response.Status != protocol.RequestResponse_OK {
					return nil, fmt.Errorf(resp.Response.Message)
				}

				return resp.Contact, nil
			},
		},
	},
}
