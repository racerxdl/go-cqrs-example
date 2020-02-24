package gql

import "github.com/graphql-go/graphql"

var schemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(rootQuery),
	Mutation: graphql.NewObject(rootMutations),
}

func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(schemaConfig)
}
