package mutation

import (
	"context"
	"ikomers-be/http/api/schema"
	"ikomers-be/model"
	"ikomers-be/model/helper"
	"ikomers-be/use_case"

	"github.com/graphql-go/graphql"
)

func NewRegisterUserField(mysqlUserRepository model.UserRepository, securityManager helper.SecurityManager) *graphql.Field {
	return &graphql.Field{
		Type:        schema.UserType,
		Description: "Register new user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"first_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"last_name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"type": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := context.Background()
			req := use_case.RegisterUserRequest{
				Email:     p.Args["email"].(string),
				Password:  p.Args["password"].(string),
				FirstName: p.Args["first_name"].(string),
				LastName:  p.Args["last_name"].(string),
				Type:      p.Args["type"].(int),
			}

			registerUserUseCase := use_case.NewRegisterUserUsecase(mysqlUserRepository, securityManager)
			res, err := registerUserUseCase.Execute(ctx, &req)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	}
}
