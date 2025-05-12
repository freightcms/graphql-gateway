package main

import (
	"github.com/freightcms/carriers/web"
	"github.com/graphql-go/graphql"
)

var (
	RootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "GatewayQueries",
		Fields: graphql.Fields{
			"carriers": &graphql.Field{
				Type: graphql.NewInputObject(web.CarrierQuery.Fields()["carriers"]),
			},
		},
	})
)
