package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type SecurityManagerMock struct {
	Mock mock.Mock
}

func (repo *SecurityManagerMock) GenerateID(ctx context.Context) (string, error) {
	arguments := repo.Mock.Called(ctx)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *SecurityManagerMock) HashPassword(ctx context.Context, password string) (string, error) {
	arguments := repo.Mock.Called(ctx, password)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *SecurityManagerMock) CompareHashAndPassword(ctx context.Context, password string, hashed string) error {
	arguments := repo.Mock.Called(ctx, password, hashed)
	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	}

	return nil
}
