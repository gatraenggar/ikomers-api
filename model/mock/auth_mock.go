package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	Mock mock.Mock
}

func (repo *AuthRepositoryMock) AddRefreshToken(ctx context.Context, refreshToken string) error {
	arguments := repo.Mock.Called(ctx, refreshToken)
	if arguments.Get(0) != nil {
		return arguments.Get(0).(error)
	}

	return nil
}
