package gql

import "github.com/graphql-go/graphql"

var contact = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Contact",
	Description: "A contact from contact list",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name:        "id",
			Description: "ID of the Contact",
			Type:        graphql.String,
		},
		"name": &graphql.Field{
			Name:        "name",
			Description: "Name of the contact",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"last_updated": &graphql.Field{
			Name:        "last_updated",
			Description: "When this contact was last updated",
			Type:        graphql.String,
		},
	},
})
