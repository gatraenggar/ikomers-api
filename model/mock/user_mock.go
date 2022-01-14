package mock

import (
	"context"
	"ikomers-be/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repo *UserRepositoryMock) GenerateID(ctx context.Context) (string, error) {
	arguments := repo.Mock.Called(ctx)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *UserRepositoryMock) HashPassword(ctx context.Context, password string) (string, error) {
	arguments := repo.Mock.Called(ctx, password)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(string), nil
	}

	return arguments.Get(0).(string), arguments.Get(1).(error)
}

func (repo *UserRepositoryMock) CheckEmailAvailability(ctx context.Context, email string) error {
	arguments := repo.Mock.Called(ctx, email)
	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

func (repo *UserRepositoryMock) RegisterUser(ctx context.Context, u *model.User) (*model.User, error) {
	arguments := repo.Mock.Called(ctx, u)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(*model.User), nil
	}

	return nil, arguments.Get(1).(error)
}