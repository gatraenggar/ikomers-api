package test

import (
	"context"
	"ikomers-be/database"
	"ikomers-be/helper/db_test"
	"ikomers-be/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthRepository(t *testing.T) {
	tableHelper := db_test.NewAuthTableTestHelper()

	newDB, err := database.NewDB()
	if err != nil {
		t.Fatalf("failed creating new db: %v", err)
	}
	authDB := repository.NewMySqlAuthRepo(newDB)

	t.Run("add refresh token", func(t *testing.T) {
		ctx := context.Background()
		refreshToken := "someRefreshTokenHere"

		authDB.AddRefreshToken(ctx, refreshToken)

		tokenRows, err := tableHelper.GetRefreshToken(refreshToken)

		assert.NoError(t, err)
		assert.Len(t, tokenRows, 1)
		tableHelper.CleanTable()
	})
}
