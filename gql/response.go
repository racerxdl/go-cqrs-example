package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/racerxdl/go-cqrs-example/protocol"
)

var requestStatusEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "RequestStatus",
	Values: graphql.EnumValueConfigMap{
		"OK": &graphql.EnumValueConfig{
			Value:       protocol.RequestResponse_OK,
			Description: "Request went OK",
		},
		"ERROR": &graphql.EnumValueConfig{
			Value:       protocol.RequestResponse_ERROR,
			Description: "Request went ERROR",
		},
	},
	Description: "Status of a Request",
})

var requestResponse = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RequestResponse",
	Description: "Response status of a request",
	Fields: graphql.Fields{
		"status": &graphql.Field{
			Type:        graphql.NewNonNull(requestStatusEnum),
			Description: "Status of the request",
		},
		"message": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Response Message",
		},
	},
})
