package test

import (
	"context"
	"ikomers-be/helper"
	"ikomers-be/model"
	"ikomers-be/model/mock"
	"ikomers-be/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

var userRepository = &mock.UserRepositoryMock{Mock: testifyMock.Mock{}}

func TestUserRepository(t *testing.T) {
	tableHelper := helper.NewUserTableTestHelper()

	t.Run("check email availability", func(t *testing.T) {
		ctx := context.Background()
		user := &model.User{
			Email:     "johndoe@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
		}

		tableHelper.AddUser(user)

		mysqlRepo := repository.NewMySqlUserRepo(tableHelper.DB)
		err := mysqlRepo.CheckEmailAvailability(ctx, user.Email)

		assert.EqualError(t, err, "email is not available")
		tableHelper.CleanTable()
	})

	t.Run("register user method", func(t *testing.T) {
		ctx := context.Background()
		user := &model.User{
			Email:     "correct@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
		}

		mockGenerateID := userRepository.Mock.On("GenerateID", ctx).Return("random-id", nil)
		user.ID = mockGenerateID.ReturnArguments.Get(0).(string)

		mysqlRepo := repository.NewMySqlUserRepo(tableHelper.DB)
		res, err := mysqlRepo.RegisterUser(ctx, user)

		rows, rowsErr := tableHelper.GetUserByID(&model.User{}, user.ID)

		assert.Equal(t, rowsErr, nil)
		assert.Len(t, rows, 1)
		assert.Equal(t, err, nil)
		assert.Equal(t, res.ID, user.ID)
		assert.Equal(t, res.Email, user.Email)
		assert.Equal(t, res.FirstName, user.FirstName)
		assert.Equal(t, res.LastName, user.LastName)
		assert.Equal(t, res.LastName, user.LastName)

		tableHelper.CleanTable()
	})
}
