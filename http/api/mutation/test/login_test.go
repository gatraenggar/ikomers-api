package test

import (
	"fmt"
	"ikomers-be/helper/db_test"
	h "ikomers-be/http"
	"ikomers-be/http/api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginMutation(t *testing.T) {
	email := "johndoe@email.com"
	password := "somePasswordHere"
	firstName := "John"
	lastName := "Doe"
	userType := 1

	var registerUserReq string = fmt.Sprintf(
		"mutation{registerUser(email:\"%s\",password:\"%s\",first_name:\"%s\",last_name:\"%s\",type:%v){email,first_name,last_name,type}}",
		email,
		password,
		firstName,
		lastName,
		userType,
	)
	h.ExecuteQuery(registerUserReq, api.NewGraphQLSchema())

	// login request:
	// `
	// 	mutation{
	// 		registerUser(
	// 			email:"johndoe@email.com",
	// 			password:"somePasswordHere",
	// 		){
	// 			first_name,
	// 			last_name,
	// 			type,
	// 			auth_token {
	// 				access_token
	//				refresh_token
	// 			}
	// 		}
	// 	}
	// `
	var loginReq string = fmt.Sprintf(
		"mutation{login(email:\"%s\",password:\"%s\"){first_name,last_name,type,auth_token{access_token,refresh_token}}}",
		email,
		password,
	)

	loginRes := h.ExecuteQuery(loginReq, api.NewGraphQLSchema())

	assert.Equal(t, false, loginRes.HasErrors(), "res.HasErrors() should return false")
	assert.Contains(t, loginRes.Data, "login")

	db_test.NewUserTableTestHelper().CleanTable()
	db_test.NewAuthTableTestHelper().CleanTable()
}
