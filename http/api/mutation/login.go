package mutation

import (
	"context"
	"ikomers-be/http/api/schema"
	"ikomers-be/model"
	"ikomers-be/model/helper"
	"ikomers-be/use_case"

	"github.com/graphql-go/graphql"
)

func NewLoginField(
	mysqlAuthRepository model.AuthRepository,
	mysqlUserRepository model.UserRepository,
	securityManager helper.SecurityManager,
	tokenManager helper.TokenManager,
) *graphql.Field {
	return &graphql.Field{
		Type:        schema.UserType,
		Description: "Login",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			ctx := context.Background()
			req := use_case.LoginRequest{
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			}

			loginUseCase := use_case.NewLoginUsecase(mysqlAuthRepository, mysqlUserRepository, securityManager, tokenManager)
			res, err := loginUseCase.Execute(ctx, &req)
			if err != nil {
				return nil, err
			}

			return res, nil
		},
	}
}
