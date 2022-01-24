package mock

import (
	"context"
	"ikomers-be/model"

	"github.com/stretchr/testify/mock"
)

type TokenManagerMock struct {
	Mock mock.Mock
}

func (repo *TokenManagerMock) GenerateAccessToken(ctx context.Context, user model.User) (string, error) {
	arguments := repo.Mock.Called(ctx, user)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *TokenManagerMock) GenerateRefreshToken(ctx context.Context, user model.User) (string, error) {
	arguments := repo.Mock.Called(ctx, user)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}
