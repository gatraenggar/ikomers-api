package model

import (
	"context"
	"errors"
	"net/mail"
)

type UserType uint

const (
	EndUser UserType = iota + 1
	Admin
)

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Password  string
	Type      UserType
}

func (u *User) ValidateFields() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return errors.New("not contain needed field")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("email is not valid")
	}

	if len(u.FirstName) < 3 || len(u.FirstName) > 15 {
		return errors.New("first name length should be 3-15")
	}

	if len(u.LastName) < 3 || len(u.LastName) > 15 {
		return errors.New("last name length should be 3-15")
	}

	if len(u.Password) < 8 || len(u.Password) > 20 {
		return errors.New("password length should be 8-20")
	}

	if u.Type != EndUser && u.Type != Admin {
		return errors.New("type should be a user type")
	}

	return nil
}

type UserRepository interface {
	CheckEmailAvailability(ctx context.Context, email string) error
	RegisterUser(ctx context.Context, user *User) (*User, error)
}
