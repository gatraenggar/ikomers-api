package test

import (
	"context"
	"ikomers-be/helper/db_test"
	"ikomers-be/helper/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurityManager(t *testing.T) {
	tableHelper := db_test.NewUserTableTestHelper()

	t.Run("generate id", func(t *testing.T) {
		ctx := context.Background()

		secManager := security.NewSecurityManager(tableHelper.DB)
		id1, err1 := secManager.GenerateID(ctx)
		id2, err2 := secManager.GenerateID(ctx)

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Len(t, id1, 36)
		assert.Len(t, id2, 36)
		assert.NotEqual(t, id1, id2)
	})
}
