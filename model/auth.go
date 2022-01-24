package model

import (
	"context"
)

type Auth struct {
	AccessToken  string
	RefreshToken string
}

type AuthRepository interface {
	AddRefreshToken(ctx context.Context, refreshToken string) error
}
