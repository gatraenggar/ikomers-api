package query

import (
	"ikomers-be/http/api/schema"

	"github.com/graphql-go/graphql"
)

func NewHelloWorldField() *graphql.Field {
	return &graphql.Field{
		Type:        schema.UserType,
		Description: "Hello world",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, isOK := p.Args["id"].(string)
			if isOK {
				return id, nil
			}

			return "Hello World!", nil
		},
	}
}
