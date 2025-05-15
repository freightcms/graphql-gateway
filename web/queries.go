package web

import (
	carriersWeb "github.com/freightcms/carriers/web"
	"github.com/graphql-go/graphql"
)

var (
	RootQuery = graphql.NewObject(q.ObjectConfig{
		Name: "GatewayQueries",
		Fields: q.Fields{
			"carriers": &qraphql.Field{
				Type: carriersWeb.CarrierObject,
				Resolve: func(p graphql.ResolveParams) (any, error) {
					return nil, nil
				},
			},
		},
	})
)
