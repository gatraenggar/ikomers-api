package mock

import (
	"context"
	"ikomers-be/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repo *UserRepositoryMock) CheckEmailAvailability(ctx context.Context, email string) error {
	arguments := repo.Mock.Called(ctx, email)
	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

func (repo *UserRepositoryMock) GetSingleUser(ctx context.Context, user *model.User) (*model.User, error) {
	arguments := repo.Mock.Called(ctx, user)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(*model.User), nil
	}

	return nil, arguments.Get(1).(error)
}

func (repo *UserRepositoryMock) RegisterUser(ctx context.Context, u *model.User) (*model.User, error) {
	arguments := repo.Mock.Called(ctx, u)
	if arguments.Get(1) == nil {
		return arguments.Get(0).(*model.User), nil
	}

	return nil, arguments.Get(1).(error)
}

func (repo *UserRepositoryMock) GetPasswordByEmail(ctx context.Context, email string) (string, error) {
	arguments := repo.Mock.Called(ctx, email)
	if arguments.Get(0) == nil {
		return "", arguments.Get(1).(error)
	}

	return arguments.Get(0).(string), nil
}
