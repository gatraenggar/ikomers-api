package api

import (
	"ikomers-be/http/api/schema"
	"ikomers-be/use_case"
	"log"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: schema.UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					_, isOK := p.Args["id"].(string)
					if isOK {
						return use_case.RegisterUserResponse{}, nil
					}
					return nil, nil
				},
			},
		},
	},
)

func NewGraphQLSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery,
		},
	)
	if err != nil {
		log.Fatalf("failed to create graphql schema %v", err)
	}

	return schema
}
