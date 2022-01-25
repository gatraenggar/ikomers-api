package model

import (
	"context"
)

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthRepository interface {
	AddRefreshToken(ctx context.Context, refreshToken string) error
}
