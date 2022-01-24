package test

import (
	"context"
	"ikomers-be/helper/security"
	"ikomers-be/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenManager(t *testing.T) {
	t.Run("generate access token", func(t *testing.T) {
		ctx := context.Background()

		tokenManager := security.NewTokenManager()
		accessToken1, err1 := tokenManager.GenerateAccessToken(ctx, model.User{
			Email:     "johndoe@gmail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "johnDoePassword",
			Type:      1,
		})
		accessToken2, err2 := tokenManager.GenerateAccessToken(ctx, model.User{
			Email:     "johndoe@gmail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "johnDoePassword",
			Type:      1,
		})

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Greater(t, len(accessToken1), 1)
		assert.Equal(t, accessToken1, accessToken2)
	})

	t.Run("generate refresh token", func(t *testing.T) {
		ctx := context.Background()

		tokenManager := security.NewTokenManager()
		refreshToken1, err1 := tokenManager.GenerateRefreshToken(ctx, model.User{
			Email:     "johndoe@gmail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "johnDoePassword",
			Type:      1,
		})
		refreshToken2, err2 := tokenManager.GenerateRefreshToken(ctx, model.User{
			Email:     "johndoe@gmail.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "johnDoePassword",
			Type:      1,
		})

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Greater(t, len(refreshToken1), 1)
		assert.Equal(t, refreshToken1, refreshToken2)
	})
}
