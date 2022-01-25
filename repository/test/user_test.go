package test

import (
	"context"
	"ikomers-be/helper/db_test"
	"ikomers-be/model"
	"ikomers-be/model/mock"
	"ikomers-be/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

var userRepository = &mock.UserRepositoryMock{Mock: testifyMock.Mock{}}

func TestUserRepository(t *testing.T) {
	userTableHelper := db_test.NewUserTableTestHelper()

	t.Run("check email availability", func(t *testing.T) {
		userTableHelper.CleanTable()
		ctx := context.Background()
		user := &model.User{
			ID:        "random-id",
			Email:     "johndoe@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
			Type:      1,
		}

		userTableHelper.AddUser(user)

		mysqlRepo := repository.NewMySqlUserRepo(userTableHelper.DB)
		err := mysqlRepo.CheckEmailAvailability(ctx, user.Email)

		assert.EqualError(t, err, "email is not available")
	})

	t.Run("get single user", func(t *testing.T) {
		userTableHelper.CleanTable()
		ctx := context.Background()
		user := &model.User{
			ID:        "random-id",
			Email:     "johndoe@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
			Type:      1,
		}

		userTableHelper.AddUser(user)

		mysqlRepo := repository.NewMySqlUserRepo(userTableHelper.DB)
		res, err := mysqlRepo.GetSingleUser(ctx, user)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, &model.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Type:      user.Type,
		}, res)
	})

	t.Run("register user", func(t *testing.T) {
		ctx := context.Background()
		user := &model.User{
			Email:     "correct@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
		}

		mockGenerateID := userRepository.Mock.On("GenerateID", ctx).Return("random-id", nil)
		userTableHelper.CleanTable()
		user.ID = mockGenerateID.ReturnArguments.Get(0).(string)

		mysqlRepo := repository.NewMySqlUserRepo(userTableHelper.DB)
		res, err := mysqlRepo.RegisterUser(ctx, user)

		rows, rowsErr := userTableHelper.GetUserByID(&model.User{}, user.ID)

		assert.Equal(t, rowsErr, nil)
		assert.Len(t, rows, 1)
		assert.Equal(t, err, nil)
		assert.Equal(t, res.ID, user.ID)
		assert.Equal(t, res.Email, user.Email)
		assert.Equal(t, res.FirstName, user.FirstName)
		assert.Equal(t, res.LastName, user.LastName)
		assert.Equal(t, res.LastName, user.LastName)
	})

	t.Run("get password by email", func(t *testing.T) {
		userTableHelper.CleanTable()
		ctx := context.Background()
		user := &model.User{
			ID:        "random-id",
			Email:     "johndoe@email.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "5h3dP4$$w012d",
			Type:      1,
		}

		userTableHelper.AddUser(user)

		mysqlRepo := repository.NewMySqlUserRepo(userTableHelper.DB)
		hashed, err := mysqlRepo.GetPasswordByEmail(ctx, user.Email)

		assert.NoError(t, err)
		assert.Equal(t, user.Password, hashed)
	})

	userTableHelper.CleanTable()
}
