package test

import (
	"ikomers-be/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFields(t *testing.T) {
	user := &model.User{}
	t.Run("not contain needed field", func(t *testing.T) {
		assert.EqualError(t, user.ValidateFields(), "not contain needed field")
	})

	user.ID = "random-id"
	user.Email = "correct@email.com"
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Password = "SomePasswordHere"

	t.Run("valid fields", func(t *testing.T) {
		assert.NoError(t, user.ValidateFields(), "valid fields should not throw error")
	})

	t.Run("email is not valid", func(t *testing.T) {
		user.Email = "wrong-email.com"
		assert.EqualError(t, user.ValidateFields(), "email is not valid")
		user.Email = "correct@email.com"
	})

	t.Run("first name length should be 3-15", func(t *testing.T) {
		user.FirstName = "De"
		assert.EqualError(t, user.ValidateFields(), "first name length should be 3-15")
		user.FirstName = "John"
	})

	t.Run("last name length should be 3-15", func(t *testing.T) {
		user.LastName = "D."
		assert.EqualError(t, user.ValidateFields(), "last name length should be 3-15")
		user.LastName = "Doe"
	})

	t.Run("password length should be 8-20", func(t *testing.T) {
		user.Password = "\\X_X/"
		assert.EqualError(t, user.ValidateFields(), "password length should be 8-20")
		user.Password = "SomePasswordHere"
	})
}
