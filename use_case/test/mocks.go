package test

import (
	"ikomers-be/model/mock"

	testifyMock "github.com/stretchr/testify/mock"
)

var AuthRepository = &mock.AuthRepositoryMock{Mock: testifyMock.Mock{}}
var UserRepository = &mock.UserRepositoryMock{Mock: testifyMock.Mock{}}
var SecurityManager = &mock.SecurityManagerMock{Mock: testifyMock.Mock{}}
var TokenManager = &mock.TokenManagerMock{Mock: testifyMock.Mock{}}
