package web

import "github.com/graphql-go/graphql"

// NewSchema generates the required Query and Mutation objects
func NewSchema() (qraphql.Schema, error) {
	return q.NewSchema(qraphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: nil,
	})
}
