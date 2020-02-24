package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

var rootMutations = graphql.ObjectConfig{
	Name: "RootMutations",
	Fields: graphql.Fields{
		"addContact": &graphql.Field{
			Type: graphql.NewNonNull(requestResponse),
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Name of the contact",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cw := protocol.ContactWriterClientFromContext(p.Context)
				if cw == nil {
					return nil, fmt.Errorf("no contact writer in context")
				}

				return cw.AddContact(p.Context, &protocol.Contact{
					Name: p.Args["name"].(string),
				})
			},
		},
		"updateContact": &graphql.Field{
			Type: graphql.NewNonNull(requestResponse),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "ID of the contact",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Description: "Name of the contact",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cw := protocol.ContactWriterClientFromContext(p.Context)
				if cw == nil {
					return nil, fmt.Errorf("no contact writer in context")
				}

				// region Check existence
				// This region might not be needed since eventual consistency does not guarantee
				// that the the user will exist in update.
				cr := protocol.ContactReaderClientFromContext(p.Context)
				if cw == nil {
					return nil, fmt.Errorf("no contact reader in context")
				}
				resp, err := cr.GetContact(p.Context, &protocol.ContactReference{
					Id: p.Args["id"].(string),
				})

				if err != nil {
					return nil, err
				}

				if resp.Response.Status != protocol.RequestResponse_OK {
					return nil, fmt.Errorf("user with id %s not found", p.Args["id"].(string))
				}
				// endregion

				return cw.UpdateContact(p.Context, &protocol.Contact{
					Id:   p.Args["id"].(string),
					Name: p.Args["name"].(string),
				})
			},
		},
		"deleteContact": &graphql.Field{
			Type: graphql.NewNonNull(requestResponse),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "ID of the contact",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				cw := protocol.ContactWriterClientFromContext(p.Context)
				if cw == nil {
					return nil, fmt.Errorf("no contact writer in context")
				}

				// region Check existence
				// This region might not be needed since eventual consistency does not guarantee
				// that the the user will exist in update.
				cr := protocol.ContactReaderClientFromContext(p.Context)
				if cw == nil {
					return nil, fmt.Errorf("no contact reader in context")
				}
				resp, err := cr.GetContact(p.Context, &protocol.ContactReference{
					Id: p.Args["id"].(string),
				})

				if err != nil {
					return nil, err
				}

				if resp.Response.Status != protocol.RequestResponse_OK {
					return nil, fmt.Errorf("user with id %s not found", p.Args["id"].(string))
				}
				// endregion

				return cw.DeleteContact(p.Context, &protocol.ContactReference{
					Id: p.Args["id"].(string),
				})
			},
		},
	},
}
