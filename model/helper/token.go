package helper

import (
	"context"
	"ikomers-be/model"
)

type Token struct {
	Manager TokenManager
}

type TokenManager interface {
	GenerateAccessToken(ctx context.Context, user model.User) (string, error)
	GenerateRefreshToken(ctx context.Context, user model.User) (string, error)
}
