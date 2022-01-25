package test

import (
	"context"
	"ikomers-be/model"
	"ikomers-be/use_case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase(t *testing.T) {
	req := &use_case.LoginRequest{
		Email:    "somebody@email.com",
		Password: "SomePasswordHere",
	}

	ctx := context.Background()
	reqUser := &model.User{Email: req.Email}

	hashedPass := UserRepository.Mock.On("GetPasswordByEmail", ctx, req.Email).Return("H@5h3dP4$$w012d", nil).ReturnArguments.Get(0).(string)

	SecurityManager.Mock.On("CompareHashAndPassword", ctx, req.Password, hashedPass).Return(nil)

	mockGetSingleUser := UserRepository.Mock.On("GetSingleUser", ctx, reqUser).Return(
		&model.User{
			Email:     req.Email,
			FirstName: "John",
			LastName:  "Doe",
			Type:      1,
		}, nil,
	).ReturnArguments.Get(0).(*model.User)

	mockGenerateAccessToken := TokenManager.Mock.On("GenerateAccessToken", ctx, *mockGetSingleUser).Return(
		"mockAccessTokenHere", nil,
	).ReturnArguments.Get(0).(string)

	mockGenerateRefreshToken := TokenManager.Mock.On("GenerateRefreshToken", ctx, *mockGetSingleUser).Return(
		"mockRefreshTokenHere", nil,
	).ReturnArguments.Get(0).(string)

	AuthRepository.Mock.On("AddRefreshToken", ctx, mockGenerateRefreshToken).Return(nil)

	LoginUseCase := use_case.NewLoginUsecase(AuthRepository, UserRepository, SecurityManager, TokenManager)

	actualRes, _ := LoginUseCase.Execute(ctx, req)
	expectedRes := &use_case.LoginResponse{
		FirstName: mockGetSingleUser.FirstName,
		LastName:  mockGetSingleUser.LastName,
		Type:      int(mockGetSingleUser.Type),
		AuthToken: model.Auth{
			AccessToken:  mockGenerateAccessToken,
			RefreshToken: mockGenerateRefreshToken,
		},
	}

	UserRepository.Mock.AssertCalled(t, "GetPasswordByEmail", ctx, req.Email)
	SecurityManager.Mock.AssertCalled(t, "CompareHashAndPassword", ctx, req.Password, hashedPass)
	UserRepository.Mock.AssertCalled(t, "GetSingleUser", ctx, reqUser)
	TokenManager.Mock.AssertCalled(t, "GenerateAccessToken", ctx, *mockGetSingleUser)
	TokenManager.Mock.AssertCalled(t, "GenerateRefreshToken", ctx, *mockGetSingleUser)
	AuthRepository.Mock.AssertCalled(t, "AddRefreshToken", ctx, mockGenerateRefreshToken)
	assert.Equal(t, expectedRes, actualRes)
}
