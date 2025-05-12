package main

import "github.com/graphql-go/graphql"

// NewSchema generates the required Query and Mutation objects
func NewSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    nil,
		Mutation: nil,
	})
}
