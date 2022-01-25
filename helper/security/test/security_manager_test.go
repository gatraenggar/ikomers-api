package test

import (
	"context"
	"ikomers-be/helper/db_test"
	"ikomers-be/helper/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurityManager(t *testing.T) {
	userTableHelper := db_test.NewUserTableTestHelper()

	t.Run("generate id", func(t *testing.T) {
		ctx := context.Background()

		secManager := security.NewSecurityManager(userTableHelper.DB)
		id1, err1 := secManager.GenerateID(ctx)
		id2, err2 := secManager.GenerateID(ctx)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Len(t, id1, 36)
		assert.Len(t, id2, 36)
		assert.NotEqual(t, id1, id2)
	})

	t.Run("hash password", func(t *testing.T) {
		ctx := context.Background()
		password := "SomePasswordHere"

		secManager := security.NewSecurityManager(userTableHelper.DB)
		hashed, err := secManager.HashPassword(ctx, password)

		assert.NoError(t, err)
		assert.NotEqual(t, hashed, password)
	})

	t.Run("compare hash and password", func(t *testing.T) {
		ctx := context.Background()

		secManager := security.NewSecurityManager(userTableHelper.DB)

		password := "SomePasswordHere"
		hashed, err := secManager.HashPassword(ctx, password)
		if err != nil {
			t.Fatalf("hash password fatal: %v", err)
		}

		err = secManager.CompareHashAndPassword(ctx, password, hashed)

		assert.NotEqual(t, password, hashed)
		assert.NoError(t, err)
	})
}
