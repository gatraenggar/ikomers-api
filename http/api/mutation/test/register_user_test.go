package test

import (
	"fmt"
	"ikomers-be/helper/db_test"
	h "ikomers-be/http"
	"ikomers-be/http/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserMutation(t *testing.T) {
	email := "johndoe@email.com"
	password := "somePasswordHere"
	firstName := "John"
	lastName := "Doe"

	// mutationReq returns:
	// `
	// 	mutation{
	// 		registerUser(
	// 			email:"johndoe@email.com",
	// 			password:"somePasswordHere",
	// 			first_name:"John",
	// 			last_name:"Doe"
	// 		){
	// 			email,
	// 			first_name,
	// 			last_name
	// 		}
	// 	}
	// `
	var mutationReq string = fmt.Sprintf(
		"mutation{registerUser(email:\"%s\",password:\"%s\",first_name:\"%s\",last_name:\"%s\"){email,first_name,last_name}}",
		email,
		password,
		firstName,
		lastName,
	)

	mutationRes := map[string]interface{}{
		"registerUser": map[string]interface{}{
			"email":      email,
			"first_name": firstName,
			"last_name":  lastName,
		},
	}

	res := h.ExecuteQuery(mutationReq, api.NewGraphQLSchema())

	assert.Equal(t, false, res.HasErrors(), "res.HasErrors() should return false")
	assert.Equal(t, mutationRes, res.Data)

	db_test.NewUserTableTestHelper().CleanTable()
}
