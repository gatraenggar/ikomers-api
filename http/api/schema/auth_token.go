package schema

import "github.com/graphql-go/graphql"

var AuthTokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthToken",
		Fields: graphql.Fields{
			"access_token": &graphql.Field{
				Type: graphql.String,
			},
			"refresh_token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
