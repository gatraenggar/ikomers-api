package test

import (
	"context"
	"ikomers-be/model"
	"ikomers-be/model/mock"
	"ikomers-be/use_case"
	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

var userRepository = &mock.UserRepositoryMock{Mock: testifyMock.Mock{}}

func TestRegisterUserUseCase(t *testing.T) {
	req := &use_case.RegisterUserRequest{
		Email:     "correct@email.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "SomePasswordHere",
	}

	ctx := context.Background()
	user := &model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}
	validateFieldsErr := user.ValidateFields()

	userRepository.Mock.On("CheckEmailAvailability", ctx, user.Email).Return(nil)

	mockGenerateID := userRepository.Mock.On("GenerateID", ctx).Return("random-id", nil)
	userID := mockGenerateID.ReturnArguments.Get(0).(string)
	user.ID = userID

	mockHashPassword := userRepository.Mock.On("HashPassword", ctx, user.Password).Return("H@5h3dP4$$w012d", nil)
	hashedPass := mockHashPassword.ReturnArguments.Get(0).(string)
	user.Password = hashedPass

	mockRegisterUser := userRepository.Mock.On("RegisterUser", ctx, user).Return(&model.User{
		ID:        userID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil).ReturnArguments.Get(0).(*model.User)

	expectedRes := &use_case.RegisterUserResponse{
		ID:        mockRegisterUser.ID,
		Email:     mockRegisterUser.Email,
		FirstName: mockRegisterUser.FirstName,
		LastName:  mockRegisterUser.LastName,
	}
	actualRes, _ := use_case.NewRegisterUserUsecase(userRepository).Execute(ctx, req)

	assert.NoError(t, validateFieldsErr, "valid fields should not throw error")
	userRepository.Mock.AssertCalled(t, "CheckEmailAvailability", ctx, req.Email)
	userRepository.Mock.AssertCalled(t, "GenerateID", ctx)
	userRepository.Mock.AssertCalled(t, "HashPassword", ctx, req.Password)
	userRepository.Mock.AssertCalled(t, "RegisterUser", ctx, user)
	assert.Equal(t, expectedRes, actualRes)
}
