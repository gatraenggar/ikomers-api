package use_case

import (
	"context"
	"ikomers-be/model"
	helperModel "ikomers-be/model/helper"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Type      int        `json:"type"`
	AuthToken model.Auth `json:"auth_token"`
}

type Login struct {
	AuthRepository  model.AuthRepository
	UserRepository  model.UserRepository
	SecurityManager helperModel.SecurityManager
	TokenManager    helperModel.TokenManager
}

func NewLoginUsecase(
	a model.AuthRepository,
	u model.UserRepository,
	s helperModel.SecurityManager,
	t helperModel.TokenManager,
) *Login {
	return &Login{
		AuthRepository:  a,
		UserRepository:  u,
		SecurityManager: s,
		TokenManager:    t,
	}
}

func (ucase *Login) Execute(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	hashed, err := ucase.UserRepository.GetPasswordByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = ucase.SecurityManager.CompareHashAndPassword(ctx, req.Password, hashed)
	if err != nil {
		return nil, err
	}

	user, err := ucase.UserRepository.GetSingleUser(ctx, &model.User{Email: req.Email})
	if err != nil {
		return nil, err
	}

	accessToken, err := ucase.TokenManager.GenerateAccessToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := ucase.TokenManager.GenerateRefreshToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	err = ucase.AuthRepository.AddRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Type:      int(user.Type),
		AuthToken: model.Auth{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
