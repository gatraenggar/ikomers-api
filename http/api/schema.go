package api

import (
	"ikomers-be/helper/db_test"
	"ikomers-be/helper/security"
	"ikomers-be/http/api/mutation"
	"ikomers-be/http/api/query"
	"ikomers-be/repository"
	"log"

	"github.com/graphql-go/graphql"
)

var gormDB = db_test.NewUserTableTestHelper().DB
var mysqlUserRepository = repository.NewMySqlUserRepo(gormDB)
var securityManager = security.NewSecurityManager(gormDB)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Query",
		Description: "Hello user",
		Fields: graphql.Fields{
			/*
			   http://localhost:8080/graphql?query={user(id:"1"){id}}
			*/
			"user": query.NewHelloWorldField(),
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/*
			http://localhost:8080/graphql?query=mutation+_{registerUser(email:"johndoe@email.com",password:"somePasswordHere",first_name:"John",last_name:"Doe",){id,email,first_name,last_name}}
		*/
		"registerUser": mutation.NewRegisterUserField(mysqlUserRepository, securityManager),
	},
})

func NewGraphQLSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
	if err != nil {
		log.Fatalf("failed to create graphql schema %v", err)
	}

	return schema
}
