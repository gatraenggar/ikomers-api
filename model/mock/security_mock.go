package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type SecurityRepositoryMock struct {
	Mock mock.Mock
}

func (repo *SecurityRepositoryMock) GenerateID(ctx context.Context) (string, error) {
	arguments := repo.Mock.Called(ctx)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *SecurityRepositoryMock) HashPassword(ctx context.Context, password string) (string, error) {
	arguments := repo.Mock.Called(ctx, password)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}
