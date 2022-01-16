package security

import (
	"context"
	"errors"
	"ikomers-be/model/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type securityManager struct {
	DB *gorm.DB
}

func NewSecurityManager(gormDB *gorm.DB) helper.SecurityManager {
	return &securityManager{
		DB: gormDB,
	}
}

func (m *securityManager) GenerateID(ctx context.Context) (string, error) {
	id := uuid.New().String()
	if len(id) != 36 {
		return "", errors.New("uuid length is not as used to be?")
	}

	return id, nil
}

func (m *securityManager) HashPassword(ctx context.Context, password string) (string, error) {
	return "", nil
}
