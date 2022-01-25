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
	userType := 1

	// register user request:
	// `
	// 	mutation{
	// 		registerUser(
	// 			email:"johndoe@email.com",
	// 			password:"somePasswordHere",
	// 			first_name:"John",
	// 			last_name:"Doe",
	// 			type: 1
	// 		){
	// 			email,
	// 			first_name,
	// 			last_name,
	// 			type,
	// 		}
	// 	}
	// `
	var mutationReq string = fmt.Sprintf(
		"mutation{registerUser(email:\"%s\",password:\"%s\",first_name:\"%s\",last_name:\"%s\",type:%v){email,first_name,last_name,type}}",
		email,
		password,
		firstName,
		lastName,
		userType,
	)

	expectedRes := map[string]interface{}{
		"registerUser": map[string]interface{}{
			"email":      email,
			"first_name": firstName,
			"last_name":  lastName,
			"type":       userType,
		},
	}
	res := h.ExecuteQuery(mutationReq, api.NewGraphQLSchema())

	assert.Equal(t, false, res.HasErrors(), "res.HasErrors() should return false")
	assert.Equal(t, expectedRes, res.Data)

	db_test.NewUserTableTestHelper().CleanTable()
}
