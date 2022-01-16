package helper

import (
	"context"
)

type Security struct {
	Manager SecurityRepository
}

type SecurityRepository interface {
	GenerateID(ctx context.Context) (string, error)
	HashPassword(ctx context.Context, password string) (string, error)
}
