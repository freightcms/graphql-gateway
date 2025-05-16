package web

import "github.com/graphql-go/graphql"

var (
	rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"carriers": &graphql.Field{
				Name: "names",
				Type: graphql.NewList(graphql.String),
			},
		},
	})
)

func NewSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}
