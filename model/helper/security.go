package helper

import (
	"context"
)

type Security struct {
	Manager SecurityManager
}

type SecurityManager interface {
	GenerateID(ctx context.Context) (string, error)
	HashPassword(ctx context.Context, password string) (string, error)
}
