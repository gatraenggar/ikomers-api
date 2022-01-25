package test

import (
	"context"
	"ikomers-be/model"
	"ikomers-be/use_case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserUseCase(t *testing.T) {
	req := &use_case.RegisterUserRequest{
		Email:     "correct@email.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "SomePasswordHere",
		Type:      1,
	}

	ctx := context.Background()
	user := &model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Type:      model.UserType(req.Type),
	}
	validateFieldsErr := user.ValidateFields()

	UserRepository.Mock.On("CheckEmailAvailability", ctx, user.Email).Return(nil)

	mockGenerateID := SecurityManager.Mock.On("GenerateID", ctx).Return("random-id", nil)
	user.ID = mockGenerateID.ReturnArguments.Get(0).(string)

	mockHashPassword := SecurityManager.Mock.On("HashPassword", ctx, user.Password).Return("H@5h3dP4$$w012d", nil)
	hashedPass := mockHashPassword.ReturnArguments.Get(0).(string)
	user.Password = hashedPass

	mockRegisterUser := UserRepository.Mock.On("RegisterUser", ctx, user).Return(&model.User{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Type:      user.Type,
	}, nil).ReturnArguments.Get(0).(*model.User)

	mockGenerateAccessToken := TokenManager.Mock.On("GenerateAccessToken", ctx, *user).Return(
		"mockAccessTokenHere", nil,
	).ReturnArguments.Get(0).(string)

	mockGenerateRefreshToken := TokenManager.Mock.On("GenerateRefreshToken", ctx, *user).Return(
		"mockRefreshTokenHere", nil,
	).ReturnArguments.Get(0).(string)

	AuthRepository.Mock.On("AddRefreshToken", ctx, mockGenerateRefreshToken).Return(nil)

	registerUserUseCase := use_case.NewRegisterUserUsecase(AuthRepository, UserRepository, SecurityManager, TokenManager)

	actualRes, _ := registerUserUseCase.Execute(ctx, req)
	expectedRes := &use_case.RegisterUserResponse{
		ID:        mockRegisterUser.ID,
		Email:     mockRegisterUser.Email,
		FirstName: mockRegisterUser.FirstName,
		LastName:  mockRegisterUser.LastName,
		Type:      int(mockRegisterUser.Type),
		AuthToken: model.Auth{
			AccessToken:  mockGenerateAccessToken,
			RefreshToken: mockGenerateRefreshToken,
		},
	}

	assert.NoError(t, validateFieldsErr, "valid fields should not throw error")
	UserRepository.Mock.AssertCalled(t, "CheckEmailAvailability", ctx, req.Email)
	SecurityManager.Mock.AssertCalled(t, "GenerateID", ctx)
	SecurityManager.Mock.AssertCalled(t, "HashPassword", ctx, req.Password)
	UserRepository.Mock.AssertCalled(t, "RegisterUser", ctx, user)
	TokenManager.Mock.AssertCalled(t, "GenerateAccessToken", ctx, *user)
	TokenManager.Mock.AssertCalled(t, "GenerateRefreshToken", ctx, *user)
	AuthRepository.Mock.AssertCalled(t, "AddRefreshToken", ctx, mockGenerateRefreshToken)
	assert.Equal(t, expectedRes, actualRes)
}
